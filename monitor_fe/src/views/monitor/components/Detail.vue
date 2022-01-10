<template>
  <el-form ref="postForm" :model="postForm" :rules="rules">
    <div class="detail-container">
      <el-col :span="15">
        <el-form-item
          label="监控名称:"
          :label-width="labelWidth"
          prop="name"
        >
          <el-input
            v-model="postForm.name"
            placeholder="监控名称"
          />
        </el-form-item>
        <el-form-item label="监控级别:" :label-width="labelWidth" prop="Mrank">
          <el-radio-group v-model="postForm.Mrank">
            <el-radio :label="1">低</el-radio>
            <el-radio :label="2">中</el-radio>
            <el-radio :label="3">高</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="监控类型:" prop="Mtype" :label-width="labelWidth">
          <el-select
            v-model="postForm.Mtype"
            placeholder="监控类型"
            @change="mTypeChange"
          >
            <el-option
              v-for="item in typeList"
              :key="item.id"
              :label="item.text"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item prop="url" label="服务url：" :label-width="labelWidth">
          <el-col :span="4">
            <el-select
              v-model="postForm.protocol"
              placeholder="协议"
            >
              <el-option label="http" value="http://"></el-option>
              <el-option label="https" value="https://"></el-option>
            </el-select>
          </el-col>
          <el-col :span="20">
            <el-input
              v-model="postForm.url"
              placeholder="url"
              style="margin-left: 2px"
            />
          </el-col>
        </el-form-item>

        <el-form-item prop="RequestHeader" label="RequestHeader:" :label-width="labelWidth">
          <el-input
            v-model="postForm.RequestHeader"
            placeholder="格式为JSON"
            type="textarea"
            :rows="2"
            @change="requestHeadersChange"
          />
        </el-form-item>
        <el-form-item prop="GetData" label="GET数据:" :label-width="labelWidth">
          <custom-alarm type="methodData"/>
          <el-input
            v-model="postForm.GetData"
            placeholder="格式为JSON"
            type="textarea"
            :rows="2"
          />
        </el-form-item>
        <el-form-item prop="PostData" label="POST数据:" :label-width="labelWidth">
          <custom-alarm type="methodData"/>
          <el-input
            v-model="postForm.PostData"
            placeholder="格式为JSON"
            type="textarea"
            :rows="2"
          />
        </el-form-item>

        <el-form-item prop="expectedCode" v-if="postForm.Mtype == 1" label="监控预期设置:" :label-width="labelWidth">
          <span class="form-item-text">预期页面返回码</span>
          <el-select v-model="postForm.expectedCode" multiple class="block">
            <el-option
              v-for="item in portList"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item v-if="postForm.Mtype == 2" label="获取接口数据" :label-width="labelWidth">
          <a
            class="el-button el-button--primary el-button--small"
            href="javascript:void(0)"
            @click="getData"
          >获取数据</a>
        </el-form-item>
        <el-form-item v-if="postForm.Mtype == 2" label="页面数据展示" :label-width="labelWidth">
          <el-input
            v-model="pageData"
            type="textarea"
            :rows="5"
          />
        </el-form-item>
        <el-form-item
          :label-width="labelWidth"
          v-if="postForm.Mtype == 2"
          prop="mRules"
          :rules="[
        { required: true, message: '请输入监控规则', trigger: 'submit' },
        { validator: validateRule, trigger: 'submit' },
      ]"
        >
          <div slot="label" style="display: inline-block">
            <span>监控规则</span>
            <custom-alarm type="rule"/>
          </div>
          <data-monitor v-model="postForm.mRules"/>
        </el-form-item>
        <el-form-item label="测试断言:" :label-width="labelWidth">
          <a
            class="el-button el-button--primary el-button--small"
            href="javascript:void(0)"
            @click="getHttpTest"
          >测试规则</a>
        </el-form-item>
        <Alarm :form="alarmForm" :alarmList="alarmList"/>
        <el-form-item :label-width="labelWidth">
          <el-row>
            <el-col :span="3">
              <el-button style="width: 120px" type="primary" @click="createMonitor">保存</el-button>
            </el-col>
            <el-col :span="3">
              <el-button style="width: 120px;margin-left: 200px" @click="goback">取消</el-button>
            </el-col>
          </el-row>
        </el-form-item>
      </el-col>
    </div>
  </el-form>
</template>

<script>
import Alarm from '@/views/monitor/components/Alarm'

const portList = [200, 301, 302, 400, 401]
import Detail from '@/views/book/components/Detail'
import DataMonitor from '@/views/monitor/components/DataMonitor'
import CustomAlarm from '@/views/monitor/components/CustomAlarm'
import {
  createMonitor,
  DataTest,
  gealarmtList,
  getData,
  getMonitorInfo,
  httpTest,
  updateMonitor
} from '../../../api/monitor'
import { monitorTypeList } from '@/views/monitor/const'
import { slice } from 'lodash/array'

const fields = {
  name: '监控项名称',
  url: '协议：url',
  expectedCode: '预期返回码',
  RequestHeader: 'RequestHeader',
  GetData: 'GetData',
  PostData: 'PostData'
}

export default {
  components: { DataMonitor, Alarm, Detail, CustomAlarm },
  props: {
    isEdit: Boolean
  },
  mounted() {
  },
  data() {
    const jsonvalidateRequire = (rule, value, callback) => {
      if (value.length === 0) {
        callback()

      } else {
        let b = this.isJSON(value)
        if (b) {
          callback()
        } else {
          callback(new Error(fields[rule.field] + '必须为json'))
        }
      }
    }

    const validateRequire = (rule, value, callback) => {
      if (value.length === 0) {
        callback(new Error(fields[rule.field] + '必须填写'))
      } else {
        callback()
      }
    }

    return {
      labelWidth: '120px',
      postForm: {
        mRules: [{ rule: 'count', arg1: '', arg2: '' }],
        protocol: 'https://',
        expectedCode: [200],
        Mrank: 1
      },
      portList,
      typeList: monitorTypeList,
      rules: {
        name: [{ validator: validateRequire }],
        expectedCode: [{ validator: validateRequire }],
        url: [{ validator: validateRequire }],
        RequestHeader: [{ validator: jsonvalidateRequire }],
        GetData: [{ validator: jsonvalidateRequire }],
        PostData: [{ validator: jsonvalidateRequire }]
      },
      pageData: '',
      alarmForm: {
        mIsdefault: 0,
        mTimeout: 60,
        mRetry: 0,
        mWaitTime: 3
      },
      alarmList: [],
      id: this.$route.params.monitorId
    }
  },
  created() {
    if (this.isEdit) {
      const monitorId = (this.$route.params.monitorId)
      this.getMonitorInfo(monitorId)
    }
    if (!this.isEdit) {
      this.postForm.Mtype = 1
      this.postForm.Mrank = 1
    }
    this.getAlarmList()
  },
  computed: {
    ruleList() {
      if (this.value && this.value.length > 0) return this.value
      return [this.$lodash.cloneDeep(this.postForm.mRules)]
    }
  },
  methods: {
    mTypeChange() {
      this.$forceUpdate()
    },
    submitForm() {

    },
    getData() {
      const getDataPost = {
        url: this.postForm.protocol + this.postForm.url,
        request_header: this.postForm.RequestHeader,
        post_data: this.postForm.PostData,
        get_data: this.postForm.GetData
      }

      getData(getDataPost).then(response => {
        this.showInterfaceData(response.data)
      }).catch(() => {

      })

    },
    getHttpTest() {
      if (this.postForm.Mtype == 1) {
        if (this.postForm.url == undefined) {
          this.$alert('请检查', '输入有误', {
            confirmButtonText: '确定'
          })
          return
        }
        const httptest = {
          url: this.postForm.protocol + this.postForm.url,
          request_header: this.postForm.RequestHeader,
          post_data: this.postForm.PostData,
          get_data: this.postForm.GetData,
          expectedCode: '[' + this.postForm.expectedCode.toString() + ']'
        }
        if (this.postForm.Mtype == 1) {
          httpTest(httptest).then(response => {
            this.showAssertResult(this.postForm.Mtype, response)
          }).catch(() => {

          })
        }
      } else if (this.postForm.Mtype == 2) {
        let b = this.checkRules(this.postForm.mRules)
        if (this.postForm.url == undefined || !b) {
          let msg = '请检查'
          if (!b) {
            msg = msg + '规则(规则不能为空)'
          } else {
            msg = msg + 'url(url不能为空)'
          }
          this.$alert(msg, '输入有误', {
            confirmButtonText: '确定'
          })
          return
        }
        const rule = this.purifyRules(this.postForm.mRules)
        const a = JSON.stringify(rule)
        const datatest = {
          url: this.postForm.protocol + this.postForm.url,
          request_header: this.postForm.RequestHeader ? this.postForm.RequestHeader : '',
          post_data: this.postForm.PostData ? this.postForm.PostData : '',
          get_data: this.postForm.GetData ? this.postForm.PostData : '',
          rules: a
        }
        DataTest(datatest).then(resopnse => {
          if (resopnse.data.isTrue == 'true') {
            this.$alert('检查通过', '校验结果', {
              confirmButtonText: '确定'
            })
          } else {
            this.$alert('检查失败:错误规则:' + resopnse.data.errorRules + ' 错误原因:' + resopnse.data.errorString, '校验结果', {
              confirmButtonText: '确定'
            })
          }
        }).catch(() => {
        })
      }

    },
    createInit() {
      if (this.isEdit) {
        this.postForm.Mtype = 1
      }
    },
    validateRule(rule, value, callback) {
      if (this.purifyRules(value).length === 0) {
        callback(new Error('请输入监控规则'))
      } else {
        callback()
      }
    },
    showAssertResult(type, resultData) {
      const data = resultData.data
      if (data === '校验成功') {
        this.$alert(data, '校验结果', {
          confirmButtonText: '确定'
        })
      } else {
        this.$alert('校验失败：' + data, '校验结果', {
          confirmButtonText: '确定'
        })
      }
    },
    requestHeadersChange() {
      // 是否有自定义的header，有自定义header后切换类型就不改变input的值了
      this.$emit('changeRequestHeader', this.form.requestHeaders)
    },

    showInterfaceData(resultData) {
      try {
        this.pageData = JSON.stringify(
          JSON.parse(resultData.body),
          undefined,
          4
        )
      } catch (e) {
        this.pageData = resultData.body
      }
    },
    purifyRules(ruleList) {
      const res = []
      ruleList.forEach((rule) => {
        const item = this.$lodash.cloneDeep(rule)
        if (
          ['count', 'equal', 'not-equal', 'include', 'gt', 'lt'].includes(item.rule) &&
          item.arg1 &&
          item.arg2
        ) {
          res.push(item)
        } else if (
          (item.rule === 'key-exist' || item.rule === 'not-null') &&
          item.arg1
        ) {
          delete item.arg2
          res.push(item)
        }
      })
      return res
    },
    checkRules(ruleList) {
      let b = true
      ruleList.forEach((rule) => {
        if (rule.rule == 'key-exist' || rule.rule == 'not-null') {
          if (rule.arg1 == '') {
            b = false
          }
        } else {
          if (rule.arg1 == '' || rule.arg2 == '') {
            b = false
          }
        }
      })
      return b
    },
    async getAlarmList() {
      const res = await gealarmtList()
      if (res.status === 0) this.alarmList = res.data
    },
    createMonitor() {

      let rule = this.postForm.mRules
      if (this.isEdit) {
        if(this.postForm.m_type == 2) {
          rule = JSON.stringify(this.postForm.mRules)
        }
        this.alarmForm.mFrequency = parseInt(this.alarmForm.mFrequency)
        this.alarmForm.mTimeout = parseInt(this.alarmForm.mTimeout)
      } else {
        rule = JSON.stringify(this.postForm.mRules)
        this.alarmForm.mFrequency = parseInt(this.alarmForm.mFrequency)
      }

      const monitor = {
        alarm_id: this.alarmForm.alarmStrategyId,
        name: this.postForm.name,
        m_rank: this.postForm.Mrank,
        m_type: this.postForm.Mtype,
        server_url: this.postForm.protocol + this.postForm.url,
        m_wait_time: this.alarmForm.mWaitTime,
        m_retry: parseInt(this.alarmForm.mRetry),
        m_isdefault: this.alarmForm.mIsdefault,
        m_cron_string: this.alarmForm.mCronString,
        request_headers: this.postForm.RequestHeader,
        post_data: this.postForm.PostData,
        get_data: this.postForm.GetData,
        m_rules: rule,
        alarm_callback_api: this.alarmForm.alarmCallbackApi,
        m_timeout: this.alarmForm.mTimeout,
        m_frequency: this.alarmForm.mFrequency,
        remarks: this.alarmForm.remarks
      }

      if (this.postForm.Mtype == 1) {
        monitor.expected_code = '[' + this.postForm.expectedCode.toString() + ']'
      }

      if (this.isEdit) {
        updateMonitor(this.id, monitor).then(response => {
          this.$router.push({
            path: `/monitor/list`
          })
        }).catch(() => {

        })
      } else {
        createMonitor(monitor).then(resopnse => {
          this.$router.push({
            path: `/monitor/list`
          })
        }).catch(() => {

        })
      }
    },
    getMonitorInfo(monitorId) {
      const params = {
        monitorId: monitorId
      }
      getMonitorInfo(params).then(response => {
        const u = response.data.server_url.split('//')
        this.postForm = {
          name: response.data.name,
          Mrank: response.data.m_rank,
          Mtype: response.data.m_type,
          protocol: u[0] + '//',
          url: u[1],
          RequestHeader: response.data.request_headers,
          PostData: response.data.post_data,
          GetData: response.data.get_data,
          mRules: response.data.m_rules
        }
        this.alarmForm = {
          alarmStrategyId: response.data.alarm_id,
          mFrequency: response.data.m_frequency,
          mTimeout: response.data.m_timeout,
          alarmCallbackApi: response.data.alarm_callback_api,
          mCronString: response.data.m_cron_string,
          mIsdefault: response.data.m_isdefault,
          mWaitTime: response.data.m_wait_time,
          mRetry: response.data.m_retry ? response.data.m_retry : 0,
          remarks: response.data.remarks
        }
        if (response.data.m_type == 1) {
          const ecode = this.ParseexpectCode(response.data.expected_code)
          this.postForm.expectedCode = ecode
        }
        if (response.data.m_type == 2) {
          this.postForm.mRules = JSON.parse(response.data.m_rules)
        }
      })
    },
    isJSON(str) {
      if (typeof str == 'string') {
        try {
          var obj = JSON.parse(str)
          if (typeof obj == 'object' && obj) {
            return true
          } else {
            return false
          }
        } catch (e) {
          return false
        }
      }
    },
    ParseexpectCode(s) {
      const c = []
      s = s.slice(1, s.length - 1)
      var split = s.split(',')
      for (var i = 0; i < split.length; i++) {
        c.push(parseInt(split[i]))
      }
      return c
    },
    goback(){
      this.$router.push({
        path: `/monitor/list`
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.detail-container {
  padding: 40px 50px 20px;

  .preview-image {
    width: 200px;
    height: 270px;
  }
}
</style>




