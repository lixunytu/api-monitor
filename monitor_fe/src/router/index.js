import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

export const constantRoutes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/auth-redirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401'),
    hidden: true,
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index'),
        name: 'Dashboard',
        meta: { title: '运维管理平台', icon: 'dashboard', affix: true }
      }
    ]
  }
]

export const asyncRoutes = [
  {
    path: '/monitor',  //路由
    name: 'monitor',
    component: Layout, //对应的组件
    meta: {title: '监控管理',icon: 'eye'}, // 定义title、icon
    redirect: '/monitor/list', //重定向 访问 monitor monitor/create
    children: [
      {
        path: '/monitor/create', //二级路由
        name: 'monitor',
        component: () => import('@/views/monitor/create'), //使用的组件
        meta: {title:"创建监控项",icon:'edit',roles:['admin']} //定义title、icon  roles 表示哪些可以访问
      },
      {
        path: '/book/edit/:fileName',
        name: '/bookEdit',
        hidden: true,
        component: () => import('@/views/book/edit'),
        meta: {title:"编辑图书",icon:'list',roles:['admin'],activeMenu:'/book/list'}
      },
      {
        path: '/monitor/edit/:monitorId',
        name: '/monitorEdit',
        hidden: true,
        component: () => import('@/views/monitor/edit'),
        meta: {title:"编辑监控项",icon:'list',roles:['admin'],activeMenu:'/monitor/list'}
      },
      {
        path: '/monitor/list',
        name: '/monitorList',
        component: () => import('@/views/monitor/list'),
        meta: {title:"监控项列表",icon:'list',roles:['admin']}
      },
    ]
  },
  {
    path: '/staff',  //路由
    name: 'staff',
    component: Layout, //对应的组件
    meta: {title: '报警管理',icon: 'documentation'}, // 定义title、icon
    redirect: '/staff/list', //重定向 访问 monitor monitor/create
    children: [
      {
        path: '/staff/list',
        name: '/staffList',
        component: () => import('@/views/staff/index'),
        meta: {title:"人员管理",icon:'list',roles:['admin'],activeMenu:'/staff/list'}
      },
      {
        path: '/alarm/list',
        name: '/alarmList',
        component: () => import('@/views/alarm/list'),
        meta: {title:"报警策略配置",icon:'edit',roles:['admin'],activeMenu:'/alarm/list'}
      },
      {
        path: '/alarm/create',
        name: '/alarmCreate',
        component: () => import('@/views/alarm/create'),
        hidden: true,
        meta: {title:"创建报警策略",icon:'edit',roles:['admin'],activeMenu:'/alarm/create'}
      },
      {
        path: '/alarm/edit/:alarmId',
        name: '/alarmEdit',
        component: () => import('@/views/alarm/edit'),
        hidden: true,
        meta: {title:"编辑报警策略",icon:'edit',roles:['admin'],activeMenu:'/alarm/edit'}
      }
    ]
  },
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
