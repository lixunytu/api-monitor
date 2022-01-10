<template>

  <el-table
    ref="tableComponent"
    v-loading="loading"
    class="table-container"
    :data="data"
    @selection-change="handleSelectionChange"
  >
    <el-table-column
      v-if="check"
      type="selection"
      width="55"
    />
    <el-table-column
      v-for="item in tableHeader"
      :key="item.id"
      :min-width="item.width"
      :label="item.label"
      :fixed="item.fixed"
      :align="item.align || 'center'"
    >
      <template slot-scope="scope">
        <!-- 文字 -->
        <div v-if="item.type === 'text'">
          <span v-if="item.prop === 'status'">{{ scope.row[item.prop] | formatStatus }}</span>
          <span v-else-if="item.prop === 'm_rank'">{{ scope.row[item.prop] | formatRank }}</span>
          <span v-else-if="item.prop === 'm_type'">{{ scope.row[item.prop] | formatType }}</span>
          <span v-else-if="item.prop === 'm_isdefault'">{{ scope.row | formatFrequency }}</span>
          <span v-else-if="item.prop === 'createdAt' || item.prop === 'lastRunTime'"
          >{{ scope.row[item.prop] | formatTimeStamp }}</span>
          <span v-else-if="item.prop === 'server_url'" :title="scope.row[item.prop]"
                v-html="formatServerUrl(scope.row)"
          />
          <span v-else>{{ scope.row[item.prop] || '-' }}</span>
        </div>
        <!-- 数组 -->
        <div v-if="item.type === 'array'">
          <el-tag v-for="tag in scope.row[item.prop]" :key="tag" class="tag">{{ tag }}</el-tag>
        </div>
        <!-- 执行机房定制 -->
        <div v-if="item.type === 'regionCustom'">
          <el-tag v-for="tag in formatMachineRegion(scope.row[item.prop])" :key="tag" class="tag">{{ tag }}</el-tag>
        </div>
        <!-- 报警定制 -->
        <div v-if="item.type === 'alarmCustom'">
          <ul class="alarm-reception">
            <li v-for="alarmType in scope.row[item.prop]" :key="alarmType">
              <span>{{ alarmType | formatAlarmType }}:&nbsp;&nbsp;</span>
              <el-tag v-for="tag in scope.row.receptionUsers[alarmType]" :key="tag.id" class="tag">{{
                  tag.name
                }}
              </el-tag>
            </li>
          </ul>
        </div>
        <!-- 链接 -->
        <div v-if="item.type === 'link'">
          <el-button type="text" @click="gotoLink(scope.row, item.link)">{{ scope.row[item.prop] }}</el-button>
        </div>
        <!-- 操作按钮 -->
        <div v-if="item.type === 'action'">
          <slot :row="scope.row"/>
        </div>
      </template>
    </el-table-column>
  </el-table>
</template>
<script>
import { parseTime } from '@/utils';
import {monitorStatusList, monitorRankList, monitorTypeList, alarmTypeList} from '@/views/monitor/const'
export default {
  name: 'TablePane',
  filters: {
    formatStatus: function(val) {
      const status = monitorStatusList.filter(item => item.id === val)
      if (status.length > 0) return status[0].text
      return val || '-'
    },
    formatRank: function(val) {
      const rank = monitorRankList.filter(item => item.id === val)
      if (rank.length > 0) return rank[0].text
      return val || '-'
    },
    formatType: function(val) {
      const type = monitorTypeList.filter(item => item.id === val)
      if (type.length > 0) return type[0].text
      return val || '-'
    },
    formatTimeStamp(timeStamp) {
      if (!timeStamp) return '-'
      return parseTime(timeStamp)
    },
    formatFrequency(data) {
      if (data.m_isdefault) { // cron表达式
        return data.m_cron_string
      } else {
        return `${parseInt(data.m_frequency)} s`
      }
    },
    formatAlarmType(alarm_type) {
      const record = alarmTypeList.filter(item => item.id === alarm_type)[0]
      return `${record.text}${record.extra}`
    }
  },
  props: {
    tableHeader: {
      type: Array,
      default: () => []
    },
    data: {
      type: Array,
      default: () => []
    },
    loading: Boolean,
    check: Boolean, // 是否需要批量勾选
    roomList: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      multipleSelection: []
    }
  },
  methods: {
    formatServerUrl(data) {
      if (data.m_type === 1 || data.m_type === 2) {
        const url = data.server_url
        return url && url.length > 50 ? `${url.substr(0, 50)}...` : url
      } else if (data.m_type === 3) {
        return `${data.mPort}@${data.ipAddress}`
      } else if (data.m_type === 4) {
        return `${data.processName}@${data.ipAddress}`
      } else if (data.m_type === 5) {
        return `${data.topicName}-${data.groupName}@${data.zookeeper}`
      } else if (data.m_type === 6) {
        return data.ulbname.replace(/,/g, '<br/>')
      } else if (data.m_type === 7) {
        return data.udbid
      }
    },
    formatMachineRegion(machineRegion) {
      if (!machineRegion) return
      const arr = []
      JSON.parse(machineRegion).map(r => {
        const room = this.roomList.find(i => i.id === r)
        room && (arr.push(room.name))
      })
      return arr
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
      this.$emit('changeSelection', val)
    },
    gotoLink(row, link) {
      this.$emit('gotoLink', row, link)
    },
  }
}
</script>

<style lang="scss" scoped>
.table-container {
  width: 100%;

  .tag {
    margin-right: 5px;
    margin-bottom: 5px;
  }

  .alarm-reception {
    list-style: none;
    margin: 0;
    padding: 0;

    li {
      margin-bottom: 8px;
    }

    li:last-child {
      margin-bottom: 0;
    }
  }
}
</style>
