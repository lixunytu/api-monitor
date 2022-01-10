<template>
  <div style="margin: 20px">
    <div style="margin: 20px">
      <el-row>
        <el-col :span="10">
          <el-input placeholder="报警策略名称" v-model="filter"/>
        </el-col>
        <el-col :span="5" style="margin-left: 10px">
          <el-button type="primary">筛选</el-button>
          <el-button>重置</el-button>
        </el-col>
      </el-row>
    </div>
    <div style="margin: 20px">
      <router-link :to="`/alarm/create`">
        <el-button class="pull-left" type="primary" style="width:120px;">新建</el-button>
      </router-link>
    </div>
    <div style="margin: 20px">
      <table-pane
        :loading="tableLoading"
        :table-header="tableHeader"
        :data="tableData"
      >
        <template v-slot:default="slotProps">
          <router-link :to="`/alarm/edit/${slotProps.row.id}`">
            <el-link type="primary" class="mr-10">编辑</el-link>
          </router-link>
          <el-link
            v-if="slotProps.row.status !== 1"
            slot="reference"
            type="danger"
            class="mr-10"
            @click="deleteItem(slotProps.row.id)"
          >删除
          </el-link>
        </template>
      </table-pane>

      <pagination
        :page="pageInfo.pageNum"
        :total="pageInfo.totalNum"
        :limit="pageInfo.pageSize"
        @pagination="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>
import { tableHeader } from './const'
import TablePane from '@/components/TablePane'
import { del, getList } from '@/api/alarm'
import Pagination from '@/components/Pagination'

export default {
  name: 'list',
  components: { Pagination, TablePane },
  data() {
    return {
      PAGE_KEY: '/alarminfo',
      filter: '',
      tableLoading: false,
      tableHeader,
      tableData: [],
      pageInfo: {
        pageNum: 1,
        pageSize: 10,
        totalNum: 0
      }
    }
  },
  created() {
    this.getAlarmList(this.pageInfo)
  },
  methods: {
    async getAlarmList(params) {
      const res = await getList(params)
      this.tableData = res.data.data
      this.pageInfo.totalNum = res.data.count
    },
    handleCurrentChange(val) {
      this.pageInfo.pageNum=val.page
      this.pageInfo.pageSize=val.limit
      this.getAlarmList(this.pageInfo)
    },
    async deleteItem(id) {
      const res = await del(id).then(response=>{
        this.$message(response.data)
        this.getAlarmList(this.pageInfo)
      })
    }
  }

}
</script>

<style scoped>
.mr-10 {
  margin-right: 10px;
}
</style>
