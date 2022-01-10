export const ruleConditionList = [
  { value: 'count', text: 'count 数量' },
  { value: 'equal', text: 'equal 全等' },
  { value: 'key-exist', text: 'key-exist 键值存在' },
  { value: 'not-null', text: 'not-null 非空' },
  { value: 'not-equal', text: 'not-equal 不等' },
  { value: 'include', text: 'include 包含' },
  { value: 'gt', text: 'gt 大于' },
  { value: 'lt', text: 'lt 小于' }
];


export const alarmTypeList = [
  { id: 1, text: '邮件', extra: '报警' },
  { id: 2, text: '钉钉', extra: '群' },
  { id: 3, text: '短信', extra: '报警' },
  { id: 4, text: '电话', extra: '报警' }
];

export const monitorTypeList = [
  { id: 1, text: 'HTTP监控' },
  { id: 2, text: '数据接口监控' },
];

export const filterList = [
  {
    id: 'id',
    placeholder: '监控项ID',
    type: 'input',
    colSpan: 6
  },
  {
    id: 'name',
    placeholder: '监控项名称',
    type: 'input',
    colSpan: 6
  },
  {
    id: 'rank',
    placeholder: '选择监控级别',
    type: 'select',
    child: [
      { id: 1, text: '高' },
      { id: 2, text: '中' },
      { id: 3, text: '低' }
    ],
    valueType: 'number',
    colSpan: 6
  },
  {
    id: 'type',
    placeholder: '选择监控项类型',
    type: 'select',
    colSpan: 6,
    valueType: 'number',
    child: monitorTypeList
  },
  {
    id: 'status',
    placeholder: '监控项状态',
    type: 'select',
    valueType: 'number',
    child: [
      { id: 0, text: '运行中' },
      { id: 1, text: '已停止' }
    ],
    colSpan: 6
  }
];

export const tableHeader = [
  { width: 50, prop: 'id', label: '序号', type: 'text', fixed: true },
  { width: 150, prop: 'name', label: '监控名称', type: 'text', fixed: true },
  { width: 70, prop: 'status', label: '状态', type: 'text', fixed: true },
  { width: 80, prop: 'm_rank', label: '监控级别', type: 'text' },
  { width: 120, prop: 'm_type', label: '监控项类型', type: 'text' },
  // { width: 120, prop: 'machineRegion', label: '执行机房', type: 'regionCustom' },
  { width: 200, prop: 'server_url', label: '服务URL', type: 'text' },
  { width: 110, prop: 'm_isdefault', label: '监控频率(s)', type: 'text' },
  { width: 160, prop: 'lastRunTime', label: '上次运行时间', type: 'text' },
  { width: 160, prop: 'createdAt', label: '创建时间', type: 'text' },
  { width: 200, label: '操作', type: 'action', fixed: 'right' }
];

export const monitorStatusList = [
  { id: 0, text: '运行中' },
  { id: 1, text: '已停止' }
];

export const monitorRankList = [
  { id: 1, text: '高' },
  { id: 2, text: '中' },
  { id: 3, text: '低' }
];
