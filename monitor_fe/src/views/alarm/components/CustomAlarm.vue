<template>
  <el-popover
    placement="right"
    width="600"
    trigger="click"
  >
    <div>
      <div v-if="type === 'mail'">
        <p>1. 自定义报警邮件标题，如果不定制则使用默认标题，默认标题如下：</p>
        <p>【夜神官方】${Machine_region} 监控发现异常，请及时处理: ${Department_name}-${Product_name}-${Monitor_name}-类别: ${M_type}-监控级别: ${M_rank} 已 ${Fail_times} 次运行失败，监控实际执行时间: ${RunTime}，响应时间: ${ResponseTime}，期待返回码: ${Expected_Code}，实际返回码: ${Status_Code}</p>
      </div>
      <div v-if="type === 'dingGroup'">
        <p>1. 自定义报警钉钉群内容，如果不定制则使用默认内容，默认内容如下：</p>
        <p>【夜神官方】${Machine_region}  监控发现异常，请及时处理: ${Department_name}-${Product_name} 名称:{Monitor_name}-类别: ${M_type}-监控级别: ${M_rank} 已 ${Fail_times} 次运行失败，监控实际执行时间: ${RunTime}，响应时间: ${ResponseTime}，期待返回码: ${Expected_Code}，实际返回码: ${Status_Code}</p>
      </div>
      <div v-if="type === 'message'">
        <p>1. 自定义报警短信内容，如果不定制则使用默认内容，默认内容如下：</p>
        <p>【夜神监控平台】业务监控${Machine_region} 发现异常，请及时处理: ${Department_name}-${Product_name}-${Monitor_name}-类别: ${M_type}-监控级别: ${M_rank} 已 ${Fail_times} 次运行失败</p>
        <p>注意：报警短信必须以<span class="danger-text">【夜神监控平台】业务监控</span>开头，否则会导致短信发送不成功</p>
      </div>
      <p>2. 占位符说明：</p>
      <el-table :data="tableData" size="small" :show-header="false">
        <el-table-column width="200" property="placeholder" />
        <el-table-column width="370" property="text" />
      </el-table>
    </div>
    <i slot="reference" class="el-icon-question question-hint" />
  </el-popover>
</template>

<script>
export default {
  name: 'CustomAlarm',
  props: {
    type: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      tableData: [
        { placeholder: '${Machine_region}', text: '执行机房' },
        { placeholder: '${Department_name}', text: '部门名称' },
        { placeholder: '${Product_name}', text: '产品线名称' },
        { placeholder: '${Monitor_name}', text: '监控项名称' },
        { placeholder: '${M_type}', text: '监控类型' },
        { placeholder: '${M_rank}', text: '监控级别' },
        { placeholder: '${Fail_times}', text: '失败次数' },
        { placeholder: '${RunTime}', text: '执行时间' },
        { placeholder: '${ResponseTime}', text: '响应时间' },
        { placeholder: '${Expected_Code}', text: '期待返回码' },
        { placeholder: '${Status_Code}', text: '实际返回码' },
        { placeholder: '${RetData}', text: '接口返回数据；${RetData.msg}代表取数据接口返回数据中msg的字段，支持多级获取' }
      ]
    };
  }
};
</script>
