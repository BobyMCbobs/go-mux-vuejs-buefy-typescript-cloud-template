export default [
  {
    path: '/',
    name: 'devices',
    component: () => import('@/frontend/views/Home.vue')
  },
  {
    path: '*',
    name: 'unknown-page',
    component: () => import('@/frontend/views/UnknownPage.vue')
  }
]
