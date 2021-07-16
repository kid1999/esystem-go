/**
* @auther: kid1999
* @date: 2021/7/16 9:04
* @desciption:  Admin
*/
<template>
  <div class="q-pa-md">
    <div class="q-gutter-y-md">
      <q-card>
        <q-tabs
          v-model="tab"
          dense
          class="text-grey"
          active-color="primary"
          indicator-color="primary"
          align="justify"
          narrow-indicator
        >
          <q-tab name="un-user" label="未审核用户"></q-tab>
          <q-tab name="un-goods" label="未审核商品" />
          <q-tab name="users" label="所有用户" />
        </q-tabs>

        <q-separator />

        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="un-user">

              <q-markup-table class="no-shadow">
                <thead>
                <tr>
                  <th class="text-left">学号</th>
                  <th class="text-right">姓名</th>
                  <th class="text-right">手机</th>
                  <th class="text-right">WX</th>
                  <th class="text-right">QQ</th>
                  <th class="text-right">申请日期</th>
                  <th class="text-right">批准</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="user in unUsers" :key="user.id">
                  <td class="text-left">{{user.StudentID}}</td>
                  <td class="text-right">{{user.Username}}</td>
                  <td class="text-right">{{user.Phone}}</td>
                  <td class="text-right">{{user.Wx}}</td>
                  <td class="text-right">{{user.QQ}}</td>
                  <td class="text-right">{{user.created_at }}</td>
                  <td class="text-right"><q-btn
                    push
                    color="secondary" icon="task_alt"
                    size="xs"
                    @click="acceptUser(user)"
                  /></td>
                </tr>
                </tbody>
              </q-markup-table>

          </q-tab-panel>

          <q-tab-panel name="un-goods">

            <q-markup-table class="no-shadow">
              <thead>
              <tr>
                <th class="text-left">商品名</th>
                <th class="text-right">描述</th>
                <th class="text-right">所有者学号</th>
                <th class="text-right">所有者学号姓名</th>
                <th class="text-right">价格</th>
                <th class="text-right">申请日期</th>
                <th class="text-right">图片</th>
                <th class="text-right">批准</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="goods in unGoods" :key="goods.id">
                <td class="text-left">{{goods.GoodsName}}</td>
                <td class="text-right">点击查看
                  <q-popup-proxy>
                    <q-banner>
                      {{goods.Description}}
                    </q-banner>
                  </q-popup-proxy>
                </td>
                <td class="text-right">{{goods.OwnerStudentID}}</td>
                <td class="text-right">{{goods.Username}}</td>
                <td class="text-right">{{goods.Price}}</td>
                <td class="text-right">{{goods.created_at}}</td>
                <td class="text-right">
                  <q-btn
                    push
                  color="teal"
                  icon="image_search"
                  size="xs"
                    @click="view(goods.ImgUrl)"
                />
                </td>

                <td class="text-right"><q-btn
                  push
                  color="secondary" icon="task_alt"
                  size="xs"
                  @click="acceptGoods(goods)"
                /></td>
              </tr>
              </tbody>
            </q-markup-table>

          </q-tab-panel>

          <q-tab-panel name="users">

            <q-markup-table class="no-shadow">
              <thead>
              <tr>
                <th class="text-left">学号</th>
                <th class="text-right">姓名</th>
                <th class="text-right">手机</th>
                <th class="text-right">WX</th>
                <th class="text-right">QQ</th>
                <th class="text-right">审核员</th>
                <th class="text-right">封闭</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="user in Users" :key="user.id">
                <td class="text-left">{{user.StudentID}}</td>
                <td class="text-right">{{user.Username}}</td>
                <td class="text-right">{{user.Phone}}</td>
                <td class="text-right">{{user.Wx}}</td>
                <td class="text-right">{{user.QQ}}</td>
                <td class="text-right">{{user.State}}</td>
                <td class="text-right"><q-btn
                  push
                  color="deep-orange" icon="highlight_off"
                  size="xs"
                  @click="closeUser(user)"
                /></td>

              </tr>
              </tbody>
            </q-markup-table>

          </q-tab-panel>
        </q-tab-panels>
      </q-card>

    </div>
  </div>

  <q-dialog v-model="imgCard">
    <q-img :src="imgUrl" />
  </q-dialog>

</template>

<script>
import { defineComponent } from 'vue'
import { deleted, get, put } from '../util/request'
import { successMsg, failMsg } from '../util/msg'

export default defineComponent({
  name: 'Admin',
  data () {
    return {
      admin: {},
      imgUrl: '',
      imgCard: false,
      tab: 'un-user',
      unUsers: [],
      unGoods: [],
      Users: []
    }
  },
  created () {
    const localUser = JSON.parse(localStorage.getItem('user'))
    if (localUser == null) {
      failMsg('请先登录！')
      this.$router.push('/login')
    }
    this.admin = localUser
    get('/admin/users').then(res => {
      this.Users = res
    })
    get('/admin/uncheckUsers').then(res => {
      this.unUsers = res
    })
    get('/admin/uncheckGoods').then(res => {
      this.unGoods = res
    })
  },
  methods: {
    view (url) {
      this.imgCard = true
      this.imgUrl = url
    },
    acceptUser (user) {
      user.State = this.admin.Username
      put('/admin/checkUser', user).then(res => {
        successMsg(user.Username + ' 审核通过')
      })
    },
    acceptGoods (goods) {
      goods.State = this.admin.Username
      put('/admin/checkGoods', goods).then(res => {
        successMsg(goods.GoodsName + ' 审核通过')
      })
    },
    closeUser (user) {
      deleted('/user', user).then(res => {
        successMsg(user.Username + ' 删除成功')
      })
    }
  }
})

</script>

<style scoped>

</style>
