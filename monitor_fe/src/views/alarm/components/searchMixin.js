/*
 * @Author: xiashan
 * @Date: 2020-09-21 17:12:19
 * @LastEditTime: 2020-12-08 17:09:29
 */
// import { API } from '@/const/const.api';
// import { publicGetPromise } from '@/api/commonRequest';
import { getList as getEmployeeList } from '@/api/staff';

export default {
  data: function() {
    return {
      searchLoading: false,
      timer: null,
      options: null
    };
  },
  methods: {
    searchChange(query) {
      this.searchLoading = true;
      if (this.timer) {
        clearTimeout(this.timer);
      }
      if (query) {
        this.timer = setTimeout(() => {
          this.remoteSearch(query);
        }, 200);
      } else {
        // 输入框中的内容被删为空时触发，此时会清除之前展示的搜索结果
        this.remoteSearch(query);
      }
    },

    async remoteSearch(query) {
      if (query !== '') {
        const params = {
          employeeKeywords: query,
          ignoreRole: true
        };
        // const res = await publicGetPromise(API.EMPLOYEE.LIST, params);
        const res = await getEmployeeList(params);
        this.options = res.data || [];
      } else {
        this.options = [];
      }
      this.searchLoading = false;
    },

    clearOption(status) {
      if (!status) {
        this.options = [];
      }
    }
  }
};
