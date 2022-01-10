<template>
  <div>
    <div class="monitorlist">
      <filter-pane
        :filter-list="filterList"
        :search-data="searchData"
        @search="onClickSearch"
        @reset="onClickReset"
      />
    </div>
    <div class="monitorlist">
      <router-link class="link-group" to="/monitor/create">
        <el-button type="primary" style="width: 100px">新建</el-button>
      </router-link>
      <el-button type="danger" style="width: 100px;margin-left: 30px" @click="deleteMonitor">批量删除</el-button>
      <el-button type="success" style="width: 100px;margin-left: 30px" @click="startMonitor">批量启用</el-button>
      <el-button type="warning" style="width: 100px;margin-left: 30px" @click="stopMonitor">批量停止</el-button>
    </div>
    <div class="monitorlist">
      <table-pane
        :loading="tableLoading"
        :table-header="tableHeader"
        :data="tableData"
        :room-list="roomList"
        :check="true"
        @changeSelection="changeSelection"
      >
        <template v-slot:default="slotProps">
          <el-link
            v-if="slotProps.row.status === 0"
            type="primary"
            class="mr-10"
            @click="inactiveItem(slotProps.row)"
          >停用
          </el-link>
          <el-link
            v-if="slotProps.row.status === 1"
            type="primary"
            class="mr-10"
            @click="activeItem(slotProps.row)"
          >启用
          </el-link>
          <router-link v-if="slotProps.row.status === 1"
                       :to="`/monitor/edit/${slotProps.row.id}`"
          >
            <el-link type="primary" class="mr-10">编辑</el-link>
          </router-link>
          <el-link
            v-if="slotProps.row.status === 1"
            type="danger"
            class="mr-10"
            @click="deleteItem(slotProps.row)"
          >删除
          </el-link>
          <router-link
            :to="`/history/index?productId=${slotProps.row.productId}&monitorId=${slotProps.row.id}&monitorName=${encodeURIComponent(slotProps.row.name)}`"
          >
            <el-link type="primary" class="mr-10">运行历史</el-link>
          </router-link>
        </template>
      </table-pane>

      <pagination
        v-show="pageInfo.totalNum>0"
        :total="pageInfo.totalNum"
        :page.sync="pageInfo.pageNum"
        :limit.sync="pageInfo.pageSize"
        @pagination="refresh"
      />
    </div>
  </div>

</template>

<script>
import Filter from '@/views/monitor/components/Filter'
import { filterList, tableHeader } from './const'
import FilterPane from '@/components/FilterPane'
import { changeMonitorInfoStatus, getMonitorList } from '@/api/monitor'
import TablePane from '@/components/TablePane'
import Pagination from '@/components/Pagination'

export default {
  name: 'list',
  components: { Pagination, TablePane, FilterPane },
  data() {
    return {
      PAGE_KEY: '/monitor',
      filterList,
      searchData: {
        id: '',
        name: '',
        rank: '',
        type: '',
        status: ''
      },
      tableLoading: true,
      tableHeader,
      roomList: [],
      multipleSelection: [],
      tableData: [],
      pageInfo: {
        pageNum: 1,
        totalNum: 0,
        pageSize: 10
      }
    }
  },
  mounted() {
    this.getMonitorList()
  },
  methods: {
    getIdList() {
      const monitorId = []
      for(var i=0;i<this.multipleSelection.length;i++) {
        monitorId.push(this.multipleSelection[i].id)
      }
      return monitorId
    },
    onClickSearch() {
      console.log(this.searchData)
    },
    onClickReset() {
      console.log('onClickReset')
    },
    deleteMonitor() {
      const monitorId = this.getIdList()
      if(monitorId.length==0) {
        this.$message.success('请选择监控项目！')
      } else {
        const params = {
          monitorId: monitorId
        }
        changeMonitorInfoStatus("del",params).then(response => {
          console.log(response)
          this.$message.success('批量删除成功!')
          this.refresh()
        })
      }
    },
    stopMonitor() {
      const monitorId = this.getIdList()
      if(monitorId.length==0) {
        this.$message.success('请选择监控项目！')
      } else {
        const params = {
          monitorId: monitorId
        }
        changeMonitorInfoStatus("stop",params).then(response => {
          console.log(response)
          this.$message.success('批量停止成功!')
          this.refresh()
        })
      }
    },
    startMonitor() {
      const monitorId = this.getIdList()
      if(monitorId.length==0) {
        this.$message.success('请选择监控项目！')
      } else {
        const params = {
          monitorId: monitorId
        }
        changeMonitorInfoStatus("start",params).then(response => {
          console.log(response)
          this.$message.success('批量启动成功!')
          this.refresh()
        })
      }
    },
    getMonitorList() {
      const params = {
        pageSize: this.pageInfo.pageSize,
        pageNum: this.pageInfo.pageNum
      }
      getMonitorList(params).then(response => {
        this.pageInfo.totalNum = response.data.count
        this.tableData = response.data.data
        this.tableLoading = false
      })
    },
    changeSelection(val) {
      this.multipleSelection = val
      console.log(val)
    },
    async activeItem(row) {
      let monitorId = [row.id]
      const params = {
        monitorId: monitorId
      }
      changeMonitorInfoStatus("start",params).then(response => {
        console.log(response)
        this.$message.success('启用成功!')
        this.refresh()
      })
    },
    async inactiveItem(row) {
      let monitorId = [row.id]
      const params = {
        monitorId: monitorId
      }
      changeMonitorInfoStatus("stop",params).then(response => {
        console.log(response)
        this.$message.success('停用成功!')
        this.refresh()
      })
    },
    async deleteItem(row) {
      let monitorId = [row.id]
      const params = {
        monitorId: monitorId
      }
      changeMonitorInfoStatus("del",params).then(response => {
        console.log(response)
        this.$message.success('删除成功!')
        this.refresh()
      })
    },

    refresh() {
      this.getMonitorList()
    },
    handleUpdate(row) {
      this.$router.push(`/monitor/edit/${row.id}`)
    },
  }
}
</script>

<style scoped>
.monitorlist {
  color: #43425D;
  margin: 20px
}

.mr-10 {
  margin-right: 10px;
}
</style>
