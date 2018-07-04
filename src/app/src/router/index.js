import Vue from 'vue'
import Router from 'vue-router'
import ListCustomers from '@/components/customers/ListCustomers'
import NewCustomer from '@/components/customers/NewCustomer'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/lista',
      name: 'Lista',
      component: ListCustomers
    }, {
      path: '/novo',
      name: 'Novo',
      component: NewCustomer
    }
  ]
})
