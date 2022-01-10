<template>
  <div
    v-loading="isLoading"
    class="container"
  >
    <div class="top-area clearfix">
      <div class="search-area">
        <el-form
          ref="searchInfo"
          :model="searchInfo"
          :inline="true"
          label-position="left"
          label-width="120px"
        >
          <el-row>
            <el-form-item label="请选择部门" prop="depart_ment_id">
              <el-select
                v-model="searchInfo.depart_ment_id"
                clearable
                filterable
                placeholder="请选择部门"
                style="width: 400px; margin-right: 20px;"
              >
                <el-option v-for="item in departmentOptions" :key="item.id" :label="item.name" :value="item.id" />
              </el-select>
            </el-form-item>
            <el-form-item label="员工信息" prop="employeeKeywords">
              <el-input
                v-model="searchInfo.employeeKeywords"
                style="width: 400px"
                placeholder="姓名 / 电话 / 邮箱"
              />
            </el-form-item>
          </el-row>
          <el-button
            type="primary"
            icon="el-icon-search"
            @click="search('searchInfo')"
          >
            查询
          </el-button>
          <el-button icon="el-icon-brush" @click="clearSearch('searchInfo')">
            清除
          </el-button>
        </el-form>
      </div>
      <div class="tabe-head-area clearfix">
        <el-button
          type="primary"
          icon="el-icon-plus"
          class="pull-left"
          @click="addEmployee"
        >
          添加员工
        </el-button>
      </div>
    </div>
    <div class="table-container">
      <el-table :data="tableData" style="width: 100%">
        <el-table-column prop="id" label="ID" show-overflow-tooltip />
        <el-table-column prop="depart_ment_id" label="所属部门" >
          <template slot-scope="scope">
            <span style="margin-left: 10px">{{ getDepartmentName(scope.row.depart_ment_id) || '--' }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="姓名" show-overflow-tooltip />
        <el-table-column
          prop="email"
          label="邮箱"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <i class="el-icon-message" />
            <span style="margin-left: 10px">{{ scope.row.email || '--' }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="phone_number"
          label="电话"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <i class="el-icon-phone" />
            <span style="margin-left: 10px">{{ scope.row.phone_number || '--' }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="ding_url"
          label="钉钉URL"
          show-overflow-tooltip
        >
          <template slot-scope="scope">
            <i class="el-icon-link" />
            <span style="margin-left: 10px">{{ scope.row.ding_url || '--' }}</span>
          </template>
        </el-table-column>
        <el-table-column fixed="right" label="操作" width="190">
          <template slot-scope="scope">
            <el-button
              size="mini"
              icon="el-icon-edit"
              @click="editEmployee(scope.row)"
            >
              编辑
            </el-button>
            <el-button
              size="mini"
              type="danger"
              icon="el-icon-delete"
              @click="editEmployee(scope.row, 'del')"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <div class="pagination-container">
      <el-pagination
        :current-page.sync="paginationConfig.currentPage"
        :page-sizes="[10, 15, 20, 30, 40, 50, 100]"
        :page-size="paginationConfig.pageSize"
        :total="paginationConfig.count"
        background
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
    <el-dialog
      :visible.sync="visible"
      :title="isEdit ? '编辑员工信息' : '添加员工信息'"
      :before-close="handleCloseModal"
      width="600px"
      class="modal-container"
    >
      <el-form
        ref="employeeInfo"
        :model="employeeInfo"
        :rules="rules"
        label-position="left"
        label-width="100px"
      >
        <el-form-item label="部门名称" prop="depart_ment_id">
          <el-select
            v-model="employeeInfo.depart_ment_id"
            placeholder="请选择部门"
            style="width: 400px"
          >
            <el-option v-for="item in departmentOptions" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>
        <el-form-item label="员工类型" prop="type">
          <el-select
            v-model="employeeInfo.type"
            placeholder="请选择类型"
            style="width: 400px"
            @change="employeeTypeChange"
          >
            <el-option label="员工" :value="0" />
            <el-option label="钉钉群" :value="1" />
          </el-select>
        </el-form-item>
        <div v-if="employeeInfo.type === 0">
          <el-form-item label="姓名" prop="name">
            <el-input
              v-model="employeeInfo.name"
              style="width: 400px"
              maxlength="200"
            />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input
              v-model="employeeInfo.email"
              style="width: 400px"
              maxlength="254"
            />
          </el-form-item>
          <el-form-item label="电话" prop="phone_number">
            <el-input
              v-model="employeeInfo.phone_number"
              style="width: 400px"
              maxlength="11"
            />
          </el-form-item>
        </div>
        <div v-else>
          <el-form-item label="钉钉群名" prop="name">
            <el-input
              v-model="employeeInfo.name"
              style="width: 400px"
              maxlength="200"
            />
          </el-form-item>
          <el-form-item label="URL" prop="ding_url">
            <el-input
              v-model="employeeInfo.ding_url"
              placeholder="自定义关键词请包含“监控”，否则会导致发消息不成功"
              style="width: 400px"
              maxlength="254"
            />
          </el-form-item>
        </div>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button
          icon="el-icon-circle-close"
          @click="closeModal('employeeInfo')"
        >取 消</el-button>
        <el-button
          type="primary"
          icon="el-icon-circle-check"
          @click="submitData('employeeInfo')"
        >
          确 定
        </el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
// import { API } from '@/const/const.api';
// import { publicGetPromise, publicPostPromise } from '@/api/commonRequest';
import { getList, create, update, del,getDepartmentList } from '@/api/staff';

export default {
  data() {
    return {
      isLoading: true,
      isEdit: false,
      visible: false,
      employeeId: null,
      userAuth: 3,
      departmentOptions: [],
      tableData: [],
      searchInfo: {
        depart_ment_id: null,
        employeeKeywords: null
      },
      employeeInfo: {
        depart_ment_id: null,
        type: 0,
        name: '',
        email: '',
        phone_number: '',
        ding_url: ''
      },
      paginationConfig: {
        pageSize: 20,
        currentPage: 1,
        count: 0
      },
      rules: {
        type: [{ required: true, message: '请选择员工类型', trigger: 'submit' }],
        depart_ment_id: [{ required: true, message: '请选择部门', trigger: 'submit' }],
        name: [{ required: true, message: '请填写名称', trigger: 'submit' }],
        email: [
          { required: true, message: '请填写邮箱', trigger: 'submit' },
          { type: 'email', message: '请输入正确的邮箱地址', trigger: ['submit'] }
        ],
        phone_number: [
          { required: true, message: '请填写手机号', trigger: 'submit' },
          { pattern: /^1[3|4|5|6|7|8][0-9]\d{8}$/, message: '请输入正确的手机号', trigger: 'submit' }
        ],
        ding_url: [
          { required: true, message: '请填写钉钉URL', trigger: 'submit' },
          { type: 'url', message: '请输入正确的钉钉URL', trigger: ['submit'] }
        ]
      }
    };
  },
  created() {
    this.employeeId = this.$store.state.user.roles[0].employeeId;
    this.initAuth();
    this.getDepartment();
    this.getEmployeeList();
  },
  methods: {
    initAuth() {
      const { roles } = this.$store.state.user;
      for (let i = 0; i < roles.length; i++) {
        const element = roles[i];
        if (element.role === 1) {
          this.userAuth = 1;
          break;
        } else if (element.role === 2) {
          this.userAuth = 2;
          break;
        } else {
          this.userAuth = 3;
          break;
        }
      }
    },
    async getDepartment() {
      const res = await getDepartmentList();
      this.departmentOptions = res.data;
    },
    async getEmployeeList(currentPage = 0) {
      this.isLoading = true;
      if (currentPage === 1) {
        this.paginationConfig.currentPage = 1;
      }
      const params = {
        depart_ment_id: this.searchInfo.depart_ment_id,
        employeeKeywords: this.searchInfo.employeeKeywords,
        pageSize: this.paginationConfig.pageSize,
        currentPage: this.paginationConfig.currentPage
      };
      if (params.depart_ment_id === null) delete params.depart_ment_id;
      if (params.employeeKeywords === null) delete params.employeeKeywords;
      try {
        const res = await getList(params);
        const employeeList = res.data && res.data.data || [];
        this.tableData = this.$lodash.cloneDeep(employeeList);
        this.paginationConfig.count = res.data.count || 0;
      } catch (error) {
        this.$message({
          showClose: true,
          message: '服务器错误，请重试或者联系开发人员。',
          type: 'error'
        });
      }
      this.isLoading = false;
    },
    addEmployee() {
      this.isEdit = false;
      this.visible = true;
    },
    async editEmployee(data, type) {
      if (type === 'del') {
        // 判断是不是自己，自己的话不能删除自己
        if (this.employeeId === data.id) {
          this.$message({
            showClose: true,
            message: '您不能删除自己！',
            type: 'error'
          });
          return false;
        }
        this.$confirm('此操作将永久删除该员工, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
          .then(async() => {
            this.isLoading = true;
            const res = await del(data.id);
            const code = res.status;
            if (code !== 0) {
              this.$message({
                showClose: true,
                message: '操作失败，请重试。',
                type: 'error'
              });
              return false;
            }
            this.getEmployeeList();
            this.$message({
              type: 'success',
              message: '删除成功!'
            });
          })
          .catch(() => {
            this.$message({
              type: 'info',
              message: '已取消删除'
            });
          });
        this.isLoading = false;
        return;
      }
      this.employeeInfo.id = data.id;
      this.employeeInfo.depart_ment_id = data.depart_ment_id;
      this.employeeInfo.type = data.type;
      this.employeeInfo.name = data.name;
      this.employeeInfo.email = data.email;
      this.employeeInfo.phone_number = data.phone_number;
      this.employeeInfo.ding_url = data.ding_url;
      this.isEdit = true;
      this.visible = true;
    },
    search() {
      this.getEmployeeList(1);
    },
    clearSearch(searchInfo) {
      this.$refs[searchInfo].resetFields();
    },
    submitData(employeeInfo) {
      this.$refs[employeeInfo].validate(async valid => {
        if (valid) {
          this.isLoading = true;
          const params = this.$lodash.cloneDeep(this.employeeInfo);
          try {
            if (this.isEdit) {
              await update(params.id,params);
            } else {
              await create(params);
            }
            this.isLoading = false;
            this.closeModal(employeeInfo);
            if (this.isEdit) {
              this.getEmployeeList();
            } else {
              this.getEmployeeList(1);
            }
          } catch (e) {
            console.log(e)
            this.isLoading = false;
          }
        }
      });
    },
    closeModal(employeeInfo) {
      this.visible = false;
      this.$refs[employeeInfo].resetFields();
      this.employeeInfo = {
        depart_ment_id: null,
        type: 0,
        name: '',
        email: '',
        phone_number: '',
        ding_url: ''
      };
    },
    handleCloseModal() {
      this.$confirm('检测到您填写的数据还未保存，确认关闭吗？')
        .then(_ => {
          this.closeModal('employeeInfo');
        })
        .catch(_ => {});
    },
    handleSizeChange(val) {
      this.paginationConfig.pageSize = val;
      this.paginationConfig.currentPage = 1;
      this.getEmployeeList();
    },
    handleCurrentChange(val) {
      this.paginationConfig.currentPage = val;
      this.getEmployeeList();
    },
    employeeTypeChange() {
      this.$refs['employeeInfo'].clearValidate();
    },
    getDepartmentName(id) {
      for(var i=0;i<this.departmentOptions.length;i++){
        if(id==this.departmentOptions[i].id) {
          return this.departmentOptions[i].name
        }
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.container {
  width: 100%;
  .top-area {
    .search-area {
      padding: 10px 0 10px 0;
      background: #ffffff;
      border-radius: 6px;
      margin: 20px
    }
    .tabe-head-area {
      padding: 10px 0 10px 0;
      background: #ffffff;
      border-radius: 6px;
      margin: 20px;
    }
  }
  .table-container {
    margin-bottom: 30px;
    margin: 20px
  }
  .pagination-container {
    text-align: center;
    padding-bottom: 30px;
    margin: 20px
  }
}
</style>
