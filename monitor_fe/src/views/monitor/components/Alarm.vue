<template>
  <div>
    <el-form ref="form" :model="form" :rules="rules">
      <el-form-item
        label="报警回调API"
        prop="alarmCallbackApi"
        :label-width="labelWidth"
      >
        <el-input v-model="form.alarmCallbackApi" class="block"/>
      </el-form-item>
      <el-form-item
        label="报警超时时间"
        prop="mTimeout"
        :label-width="labelWidth"
      >
        <el-row>
          <el-col :span="12">
            <el-input v-model="form.mTimeout" class="form-input-embed-m"/>
          </el-col>
          <el-col :span="12">
            <span class="form-item-text" style="margin-left: 3px">秒 (s)</span>
          </el-col>
        </el-row>
      </el-form-item>
      <el-form-item
        label="重试次数"
        prop="mRetry"
        :label-width="labelWidth"
      >
        <el-col :span="12">
          <el-input
            v-model="form.mRetry"
            class="form-input-embed-m"
            type="number"
            max="3"
            min="0"
          />
        </el-col>
      </el-form-item>
      <el-form-item
        label="重试间隔时间"
        prop="mWaitTime"
        :label-width="labelWidth"
      >
        <el-row>
          <el-col :span="12">
            <el-input
              v-model="form.mWaitTime"
              class="form-input-embed-m"
              type="number"
              max="20"
              min="1"
            />
          </el-col>
          <el-col :span="11" style="margin-left: 3px">
            <span class="form-item-text">秒 (s)</span>
          </el-col>
        </el-row>
      </el-form-item>
      <el-form-item
        label="监控执行频率"
        prop="mFrequency"
        :label-width="labelWidth"
      >
        <div slot="label" class="inline-block">
          <span>监控执行频率</span>
          <custom-alarm v-if="form.mIsdefault == 0" type="frequency"/>
          <custom-alarm v-if="form.mIsdefault == 1" type="cron-frequency"/>
        </div>
        <el-radio v-model="form.mIsdefault" :label="0">默认</el-radio>
        <el-radio v-model="form.mIsdefault" :label="1">自定义（cron表达式）</el-radio>
        <el-form-item
          prop="mFrequency"
          :rules="[{ validator: validateRequire }]"
        >
          <el-row v-if="form.mIsdefault == 0">
            <el-row>
              <el-col :span="12">
                <el-input v-model="form.mFrequency" class="form-input-embed-m" type="number" min="0"/>
              </el-col>
              <el-col :span="11" style="margin-left: 3px">
                <span class="form-item-text">秒 (s)</span>
              </el-col>
            </el-row>
          </el-row>
        </el-form-item>
        <el-row v-if="form.mIsdefault == 1">
          <el-form-item
            prop="mCronString"
            :rules="[{ validator: validateRequire }]"
          >
            <el-input v-model="form.mCronString" class="form-input-embed-l">
              <a
                slot="append"
                plain
                class="el-button el-button--default el-button--small"
                href="javascript:void(0)"
                @click="parseCrontab(form.mCronString)"
              >解析</a>
            </el-input>
          </el-form-item>
          <el-popover
            v-if="crontab && cronResult.length"
            placement="top-start"
            :title="crontab"
            width="200"
            trigger="hover"
          >
            <p>接下来10次执行结果</p>
            <p v-for="(time, index) in cronResult" :key="index">{{ time }}</p>
            <el-link slot="reference" type="primary">查看解析结果</el-link>
          </el-popover>
        </el-row>
      </el-form-item>
      <el-form-item
        label="报警策略"
        prop="alarmStrategyId"
        :label-width="labelWidth"
      >
        <el-select
          v-model="form.alarmStrategyId"
          placeholder="请选择报警策略"
          class="block"
          :popper-append-to-body="false"
        >
          <el-option
            v-for="item in alarmList"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="备注"
        prop="remarks"
        :label-width="labelWidth"
      >
        <el-input v-model="form.remarks" class="block"/>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import CustomAlarm from '@/views/monitor/components/CustomAlarm'
import { parseCron, gealarmtList } from '@/api/monitor'

const fields = {
  mTimeout: '超时时间',
  mRetry:"重试次数",
  mWaitTime:"重试间隔时间",
  alarmStrategyId:"报警策略",
  mCronString:"cron表达式",
  mFrequency:"频率",

}

export default {
  components: { CustomAlarm },
  name: 'Alarm',
  props: {
    form: {
      default() {
        return {}
      }
    },
    alarmList: {
      type: Array,
      default() {
        return []
      }
    }
  },
  data() {
    const validateRequire = (rule, value, callback) => {
      if (value.length === 0) {
        callback(new Error(fields[rule.field] + '必须填写'))
      } else {
        callback()
      }
    }
    return {
      crontab: '',
      cronResult: [],
      labelWidth: '120px',
      rules: {
        mTimeout: [{ validator: validateRequire }],
        mRetry: [{ validator: validateRequire }],
        mWaitTime: [{ validator: validateRequire }],
        alarmStrategyId: [{ validator: validateRequire }],
      },
    }
  },
  methods: {
    async parseCrontab(exp) {
      if (!exp) return
      try {
        const res = await parseCron(exp)
        this.crontab = exp
        this.cronResult = res.data
      } catch (e) {
        this.$message.error('解析失败，请检查表达式!')
      }
    },
    validateRequire(rule, value, callback){
      console.log(rule)
      if (value.length === 0) {
        callback(new Error(fields[rule.field] + '必须填写'))
      } else {
        callback()
      }
    }
  }
}
</script>

<style scoped>

</style>
