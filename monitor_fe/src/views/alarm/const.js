
export const realTimeList = [
  { id: 1, text: '高' },
  { id: 2, text: '中' },
  { id: 3, text: '低' }
];

export const alarmTypeList = [
  { id: 1, text: '邮件', extra: '报警' },
  { id: 2, text: '钉钉', extra: '群' },
  { id: 3, text: '短信', extra: '报警' },
  { id: 4, text: '电话', extra: '报警' }
];

export const editRules = {
  name: [{ required: true, message: '请输入报警策略名称', trigger: 'submit' }],
  alarm_type: [{ required: true, message: '请选择报警形式', trigger: 'submit' }],

  recipient_mail: [
    { required: true, message: '请选择报警邮件接收者', trigger: 'submit' }
  ],
  ding_group: [
    { required: true, message: '请选择报警钉钉群', trigger: 'submit' }
  ],
  ding_group_fail_num: [
    { required: true, message: '请输入报警的失败次数', trigger: 'submit' },
    { pattern: /^[1-3]$/, message: '报警失败次数大于1并小于等于3次', trigger: 'submit' }
  ],
  ding_group_alarm_num: [
    { required: true, message: '输入连续次数', trigger: 'submit' }
  ],
  ding_group_summary_time: [
    { required: true, message: '输入汇总时间', trigger: 'submit' }
  ],

  recipient_message: [
    { required: true, message: '请选择报警短信接收者', trigger: 'submit' }
  ],
  message_fail_num: [
    { required: true, message: '请输入短信报警的失败次数', trigger: 'submit' },
    { pattern: /^[1-3]$/, message: '报警失败次数大于1并小于等于3次', trigger: 'submit' }
  ],
  message_alarm_num: [
    { required: true, message: '输入连续次数', trigger: 'submit' }
  ],
  message_summary_time: [
    { required: true, message: '输入汇总时间', trigger: 'submit' }
  ],
  messageCombine: [
    { required: true, message: '输入报警次数和汇总时间', trigger: 'submit' }
  ],
  recipient_phone: [
    { required: true, message: '请选择报警电话接收者', trigger: 'submit' }
  ],
  phone_fail_num: [
    { required: true, message: '请输入电话报警的失败次数', trigger: 'submit' },
    { pattern: /^[1-3]$/, message: '报警失败次数大于1并小于等于3次', trigger: 'submit' }
  ],
  phone_alarm_num: [
    { required: true, message: '输入连续次数', trigger: 'submit' }
  ],
  phone_summary_time: [
    { required: true, message: '输入汇总时间', trigger: 'submit' }
  ]
};

export const tableHeader = [
  { width: 50, prop: 'id', label: '序号', type: 'text' },
  { width: 200, prop: 'name', label: '报警策略名称', type: 'text' },
  { width: 300, prop: 'alarm_type', label: '报警接收人', type: 'alarmCustom', align: 'left' },
  { width: 160, prop: 'createdAt', label: '创建时间', type: 'text' },
  { width: 200, label: '操作', type: 'action' }
];

