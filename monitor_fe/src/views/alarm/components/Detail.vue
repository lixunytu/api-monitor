<template>
  <div class="form">
    <el-page-header :content=" title+'报警策略'" class="page-title" @back="goBack"/>
    <el-form
      ref="postForm"
      :model="postForm"
      :rules="rules"
      label-width="150px"
      class="common-form"
    >
      <el-form-item
        label="报警策略名称"
        prop="name"
      >
        <el-input v-model="postForm.name" class="block" maxlength="200"/>
      </el-form-item>

      <el-form-item
        label="实时性"
      >
        <el-radio
          v-for="item in realTimeList"
          :key="item.id"
          v-model="postForm.real_time"
          :label="item.id"
        >
          {{ item.text }}
        </el-radio>
      </el-form-item>

      <el-form-item
        label="报警形式"
        prop="alarmType"
      >
        <el-checkbox-group v-model="postForm.alarm_type">
          <el-checkbox
            v-for="item in alarmTypeList"
            :key="item.id"
            :label="item.id"
          >{{ item.text }}
          </el-checkbox>
        </el-checkbox-group>
      </el-form-item>

      <alarm-ding v-if="showMail || showDing " :form="postForm" :show-mail="showMail" :show-ding="showDing" :is-edit="showIsEdit"/>
      <alarm-message v-if="showMessage" :form="postForm"/>
      <alarm-phone v-if="showPhone" :form="postForm"/>

      <el-form-item label="备注" prop="remarks">
        <el-input v-model="postForm.remarks" class="block" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" style="width: 120px;" @click="submitForm('postForm')">保存</el-button>
        <el-button style="width: 120px;" @click="goBack">取消</el-button>
      </el-form-item>

    </el-form>
  </div>
</template>

<script>
import AlarmDing from '@/views/alarm/components/AlarmDing'
let alarmPath = ''
import { realTimeList,alarmTypeList,editRules } from '../const'
import AlarmMessage from '@/views/alarm/components/AlarmMessage'
import AlarmPhone from '@/views/alarm/components/AlarmPhone'
import {create,getListForId,update} from '@/api/alarm'

export default {
  name: 'Detail',
  components: { AlarmPhone, AlarmMessage, AlarmDing },
  props: {
    isEdit: Boolean
  },
  data() {
    return {
      title: '新建',
      id: this.$route.params.alarmId,
      postForm: {
        real_time: 1,
        alarm_type: [1,2],
        ding_group_fail_num: 2,
        ding_group_alarm_num: 10,
        ding_group_summary_time: 60,
        message_fail_num: 2,
        message_alarm_num: 10,
        message_summary_time: 60,
        message_off_day_alarm: 0,
        phone_fail_num: 2,
        phone_alarm_num: 10,
        phone_summary_time: 60,
        phone_off_day_alarm: 0
      },
      realTimeList,
      alarmTypeList,
      rules: editRules
    }
  },
  computed: {
    showMail() {
      return this.postForm.alarm_type.includes(1);
    },
    showDing() {
      return this.postForm.alarm_type.includes(2);
    },
    showMessage() {
      return this.postForm.alarm_type.includes(3);
    },
    showPhone() {
      return this.postForm.alarm_type.includes(4);
    },
    showIsEdit() {
      return this.isEdit
    }
  },
  beforeRouteEnter(to, from, next) {
    if (!from.name) {
      alarmPath = '/alarm/index'
    } else {
      alarmPath = ''
    }
    next()
  },

  created() {
    if(this.isEdit) {
      this.title="编辑"
      const alarmId = (this.$route.params.alarmId)
      getListForId(alarmId).then(response =>{
        this.postForm = response.data
        if(this.showDing) {
          this.postForm.ding_group=this.ParseexpectCode(this.postForm.ding_group)
        }
        if(this.showMail) {
          this.postForm.recipient_mail=this.ParseexpectCode(this.postForm.recipient_mail)
        }
        if(this.showPhone) {
          this.postForm.recipient_phone=this.ParseexpectCode(this.postForm.recipient_phone)
        }
        if(this.showMessage) {
          this.postForm.recipient_message=this.ParseexpectCode(this.postForm.recipient_message)
        }
      })
    }
  },

  methods: {
    ParseexpectCode(s) {
      const c = []
      s = s.slice(1, s.length - 1)
      var split = s.split(',')
      for (var i = 0; i < split.length; i++) {
        c.push(parseInt(split[i]))
      }
      return c
    },
    goBack() {
      if (alarmPath) {
        this.$router.push(alarmPath)
      } else {
        this.$router.back()
      }
    },
    submitForm(formName) {
      // console.log(this.options)
      this.$refs[formName].validate(async valid => {
        if (valid) {
          await this.verifySuccess()
        } else {
          return false;
        }
      });
    },
    async verifySuccess() {
      const params = this.postForm
      this.loading = true;
      try {
        let res;
        if (this.isEdit) {
          params.alarm_type = "["+this.postForm.alarm_type.toString()+"]"
          if (this.showDing) {
            params.ding_group = "["+this.postForm.ding_group.toString()+"]"
          }
          if(this.showMail){
            params.recipient_mail = "["+this.postForm.recipient_mail.toString()+"]"
          }
          if(this.showMessage) {
            params.recipient_message = "["+this.postForm.recipient_message.toString()+"]"
          }
          if(this.showPhone) {
            params.recipient_phone = "["+this.postForm.recipient_phone.toString()+"]"
          }
          res = await update(this.id, params);
        } else {
          params.alarm_type = "["+this.postForm.alarm_type.toString()+"]"
          if (this.showDing) {
            params.ding_group = "["+this.postForm.ding_group.toString()+"]"
          }
          if(this.showMail){
            params.recipient_mail = "["+this.postForm.recipient_mail.toString()+"]"
          }
          if(this.showMessage) {
            params.recipient_message = "["+this.postForm.recipient_message.toString()+"]"
          }
          if(this.showPhone) {
            params.recipient_phone = "["+this.postForm.recipient_phone.toString()+"]"
          }
          if(this.isEdit) {

          }else {
            res = await create(params);
          }
        }
        this.loading = false;
        if (res.status === 0) {
          this.$message.success('保存成功');
          this.goBack();
        }
      } catch (e) {
        this.loading = false;
      }
    },
  }
}
</script>

<style scoped>
.form {
  margin: 20px;
}

.common-form {
  width: 80%;
  margin: 32px auto;
}

.page-title {
  color: #43425D;
  margin-bottom: 16px;
  font-size: 20px;
}
</style>
