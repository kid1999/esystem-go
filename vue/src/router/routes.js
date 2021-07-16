
const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/Index.vue') },
      { path: '/login', component: () => import('pages/Login') },
      { path: '/register', component: () => import('pages/Register') },
      { path: '/user', component: () => import('pages/UserInfo') },
      { path: '/addGoods', component: () => import('pages/AddGoods') },
      { path: '/admin', component: () => import('pages/Admin') }
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
