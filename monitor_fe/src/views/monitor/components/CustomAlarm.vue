<!--
 * @Author: xiashan
 * @Date: 2020-11-21 16:15:46
 * @LastEditTime: 2021-01-28 16:36:44
-->
<template>
  <el-popover
    placement="right"
    width="600"
    trigger="click"
  >
    <div v-if="type === 'methodData'">
      <p>支持以下获取动态数据的宏：</p>
      <p>1. <b>${timestamp}</b> 获取当前时间戳(到现在的秒数)</p>
      <p>2. <b>${md5NoxMobile}</b> 计算签名，待签名数据的组成规则：'nox_video'+timestamp+url+'nox_video'；参数中需要设置timestamp和url</p>
      <p>3. <b>${getRandomStr}</b> 'ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz0123456789' 中随机取16个字符拼接成字符串</p>
      <div><span class="info">eg.</span><br>{</div>
      <div style="padding-left: 10px;">"url": "https://www.youtube.com/watch?v=oGRWd1nTNhs",<br>
        "timestamp": "${timestamp}",<br>
        "sign": "${md5NoxMobile}"</div>
      <div>}</div>
      <div v-if="page === 'BatchNew'">
        <hr>
        <p>批量生成操作支持以下2个函数：</p>
        <p>1. <b>${forEach([array])}</b>循环数组的每一个元素</p>
        <p>
          <span class="info">参数</span><br>
          array：forEach()函数调用的数组，只支持一维数组<br>
          <span class="info">eg.</span><br>
          ${forEach(['JP','TW', 'KR'])}
        </p>
        <p>2. <b>${getValue(obj, key)}</b>根据键获取对象中的值</p>
        <p>
          <span class="info">参数</span><br>
          obj：获取键值的对象<br>
          key：获取obj中key值对应的value(一般为对象中其他字段)<br>
          <span class="info">eg.</span><br>
          ${getValue({<br>
          &nbsp;&nbsp;'JP': ['com.global.wfwd.google','com.ulugame.swordkingdoms.google'],<br>
          &nbsp;&nbsp;'TW': ['jp.glee.girl','jp.co.ultimedia.projectsteam','jp.kkdw.kenkadou'],<br>
          }, region)}
        </p>
      </div>
    </div>
    <div v-if="type === 'cron-frequency'">
      <p>1. 结构</p>
      <p>corn从左到右（用空格隔开）：秒 分 小时 月份中的日期 月份 星期中的日期</p>
      <p>2. 各字段的含义</p>
      <el-table :data="tableData" size="small">
        <el-table-column label="字段" property="field" />
        <el-table-column label="允许值" property="desc" />
        <el-table-column label="允许的特殊字符" property="remarks" />
      </el-table>
      <p>3. 常用表达式例子</p>
      <p>
        0 0/3 * * * ?   表示每3分钟执行任务<br>
        0 0 0/1 * * ?   表示每小时执行任务<br>
        0 0 12 * * ?   表示每天中午12点触发<br>
        0 15 10 ? * MON-FRI   表示周一到周五每天上午10:15执行任务
      </p>
    </div>
    <div v-if="type === 'frequency'">
      <p>仅支持1 2 3 4 5 6 10 15 20 30 60  120 180 240  360  480  720  1440 mins 输入</p>
    </div>
    <div v-if="type === 'rule'">
      <p>监控规则使用说明</p>
      <p>
        <el-link
          type="primary"
          target="_blank"
          :underline="false"
          href="http://team.yeshen.com/pages/viewpage.action?pageId=37492517"
        >
          http://team.yeshen.com/pages/viewpage.action?pageId=37492517
        </el-link>
      </p>
    </div>
    <div v-if="type === 'ip'">
      <p>可以填写多个IP地址或者域名</p>
      <p>如果填写了IP地址或者域名，服务URL只需要填写路径即可；<br>如果没有填写IP地址或者域名，服务URL需要填写完整的URL</p>
    </div>
    <div v-if="type === 'ucloudName'">
      <p>绿色字体表示配置过监控项，可使用全局搜索进行查询</p>
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
    },
    page: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      tableData: [
        { field: '秒（Seconds)', desc: '0~59的整数', remarks: ', - * / 四个字符' },
        { field: '分（Minutes）', desc: '0~59的整数', remarks: ', - * / 四个字符' },
        { field: '小时（Hours)', desc: '0~23的整数', remarks: ', - * / 四个字符' },
        { field: '日期（DayofMonth)', desc: '1~31的整数（但是你需要考虑你月的天数）', remarks: ',- * ? / L W C 八个字符' },
        { field: '月份（Month）', desc: '1~12的整数或者 JAN-DEC', remarks: ', - * / 四个字符' },
        { field: '星期（DayofWeek）', desc: '1~7的整数或者 SUN-SAT （1=SUN）', remarks: ' , - * ? / L C # 八个字符' }
      ]
    };
  }
};
</script>

<style lang="scss" scoped>
.info {
  background-color: #efefef;
  padding: 0 8px;
  border-radius: 6px;
}
</style>
