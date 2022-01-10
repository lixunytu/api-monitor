<template>
  <div class="alarm-wrap">
    <el-form-item v-if="showMail" label="报警邮件接收者" prop="recipient_mail">

      <el-select
        v-model="form.recipient_mail"
        class="block"
        multiple
        filterable
        remote
        placeholder="请选择报警邮件接收者"
        :reserve-keyword="false"
        :remote-method="searchChange"
        :loading="searchLoading"
        @visible-change="clearOption"
      >
        <el-option
          v-for="item in mailOptions"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        >
          <span>{{ item.name }} - {{ item.email }}</span>
        </el-option>
      </el-select>

    </el-form-item>

    <el-form-item v-if="showDing" label="报警钉钉群" prop="ding_group">

      <el-select
        v-model="form.ding_group"
        class="block"
        multiple
        filterable
        remote
        placeholder="请选择报警钉钉群"
        :reserve-keyword="false"
        :remote-method="searchChange"
        :loading="searchLoading"
        @visible-change="clearOption"
      >
        <el-option
          v-for="item in dingOptions"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        >
          <span>{{ item.name }}</span>
        </el-option>
      </el-select>

    </el-form-item>

    <el-form-item label="报警频率" prop="ding_group_fail_num">
      连续<el-input v-model="form.ding_group_fail_num" class="form-input-embed" />次失败后，启动报警
    </el-form-item>

    <el-form-item label="报警策略" class="is-required">
      连续报警
      <el-form-item class="inline-block" prop="ding_group_alarm_num">
        <el-input v-model="form.ding_group_alarm_num" class="form-input-embed" />
      </el-form-item>
      次，
      后续
      <el-form-item class="inline-block" prop="ding_group_summary_time">
        <el-input v-model="form.ding_group_summary_time" class="form-input-embed" />
      </el-form-item>
      分钟汇总
    </el-form-item>

  </div>
</template>

<script>
import searchMixin from './searchMixin';
import CustomAlarm from './CustomAlarm';
import { slice } from 'lodash/array'

export default {
  name: 'AlarmDing',
  components: { CustomAlarm },
  mixins: [searchMixin],
  props: {
    form: {
      type: Object,
      default() {
        return {};
      }
    },
    showMail: {
      type: Boolean,
      default: true
    },
    showDing: {
      type: Boolean,
      default: true
    },
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    mailOptions() {
      if (this.options === null) {
        return this.form.receptionUsers && this.form.receptionUsers[1];
      } else {
        return this.options;
      }
    },
    dingOptions() {
      if (this.options === null) {
        return this.form.receptionUsers && this.form.receptionUsers[2];
      } else {
        return this.options;
      }
    }
  },
  created() {
  },
  methods:{
  }
}
</script>

<style scoped>
.inline-block {
  display: inline-block;
}

.alarm-wrap {
  padding-top: 12px;
  background-color: #fafafa;
}

.form-input-embed {
  width: 60px;
  margin: 0 8px;
}
</style>
