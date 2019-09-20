import Vue from 'vue'
import Router from 'vue-router'
// import Home from '@/views/Home.vue'
import layout from '@/layout'

Vue.use(Router);

export default new Router({
  routes: [
    // {
    //   path: '/',
    //   component: layout,
    //   redirect: '/dashboard',
    //   children: [
    //     {
    //       path: 'dashboard',
    //       name: 'dashboard',
    //       component: ()=>import('@/views/dashboard/index'),
    //       meta: {title:"dashboard", icon:"dashboard", affix:true},
    //     }
    //   ]
    // },
    {
      path: '/',
      // component: layout,
      redirect: '/document/index',
    },
    {
      path: '/document',
      component: layout,
      children: [
        {
          path: 'index',
          name: 'document',
          component: ()=>import('@/views/document/index'),
          meta: {title:"document", icon:"document", affix:true},
        }
      ]
    },
    {
      path: '/mock',
      component: layout,
      redirect: '/mock/index',
      children: [
        {
          path: 'index',
          name: 'mock',
          component: ()=>import('@/views/mock/index'),
          meta: {title:"mock", icon:"mock", affix:true},
        }
      ]
    }
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (about.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    // }
  ]
})
