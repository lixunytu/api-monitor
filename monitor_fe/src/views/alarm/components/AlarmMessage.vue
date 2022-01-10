<template>
  <div class="alarm-wrap">
    <el-form-item label="报警短信接收者" prop="recipient_message">

      <el-select
        v-model="form.recipient_message"
        class="block"
        multiple
        filterable
        remote
        placeholder="请选择报警短信接收者"
        :reserve-keyword="false"
        :remote-method="searchChange"
        :loading="searchLoading"
        @blur="clearOption"
      >
        <el-option
          v-for="item in messageOptions"
          :key="item.id"
          :label="item.name"
          :value="item.id"
        >
          <span>{{ item.name }} - {{ item.phoneNumber }}</span>
        </el-option>
      </el-select>

    </el-form-item>

    <el-form-item label="短信报警频率" prop="message_fail_num">
      连续<el-input v-model="form.message_fail_num" class="form-input-embed" />次失败后，启动报警
    </el-form-item>

    <el-form-item label="短信报警策略" class="is-required">
      连续报警
      <el-form-item class="inline-block" prop="message_alarm_num">
        <el-input v-model="form.message_alarm_num" class="form-input-embed" />
      </el-form-item>
      次，
      后续
      <el-form-item class="inline-block" prop="message_summary_time">
        <el-input v-model="form.message_summary_time" class="form-input-embed" />
      </el-form-item>
      分钟汇总
    </el-form-item>

    <el-form-item label="短信报警例外">
      <el-checkbox v-model="form.message_off_day_alarm" true-label="1" false-label="0" :checked="form.message_off_day_alarm == 1">低、中级别报警非工作时间不发短信（高级别报警照常发送）</el-checkbox>
    </el-form-item>
  </div>
</template>

<script>
import searchMixin from './searchMixin';

export default {
  name: 'AlarmMessage',
  mixins: [searchMixin],
  props: {
    form: {
      type: Object,
      default() {
        return {};
      }
    }
  },
  computed: {
    messageCombine() {
      return `${this.form.messageAlarmNum}${this.form.messageSummaryTime}`;
    },
    messageOptions() {
      if (this.options === null) {
        return this.form.receptionUsers &&
          this.form.receptionUsers[3] &&
          this.form.receptionUsers[3].map(item => {
            item.phoneNumber = item.phone_number;
            return item;
          });
      } else {
        return this.options;
      }
    }
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
