<template>

  <div class="row">
    <div class="col-xs-12 col-sm-6 col-md-4">
      <q-card class="my-card" flat bordered>
        <q-img
          src="https://cdn.quasar.dev/img/parallax2.jpg">

          <div class="absolute-bottom text-subtitle text-center">
            18.8
            元
          </div>
        </q-img>
        <q-card-section>
          <div class="text-overline text-orange-9">拥有者</div>
          <div class="text-h5 q-mt-sm q-mb-xs">商品名称</div>
          <div class="text-caption text-grey">
            描述信息
          </div>
        </q-card-section>

        <q-card-actions>
          <q-btn flat color="dark" label="联系方式" @click="ViewPhone"/>
          <q-btn flat color="dark" label="msg" @click="Test"/>
          <q-btn flat color="dark" label="TestAxios" @click="TestAxios"/>
        </q-card-actions>

      </q-card>
    </div>

  </div>

  <q-dialog v-model="view">
    <q-card style="width: 300px" class="q-px-sm q-pb-md">
      <q-card-section>
        <div class="text-h6">联系方式</div>
      </q-card-section>

      <q-item-label header>手机</q-item-label>
      <q-item dense>
        <q-item-section avatar>
          <q-icon name="phone_iphone" />
        </q-item-section>
        <q-item-section>
          hello
        </q-item-section>
      </q-item>

      <q-item-label header>QQ</q-item-label>
      <q-item dense>
        <q-item-section avatar>
          <q-icon name="sports_kabaddi" />
        </q-item-section>
        <q-item-section>
          1234565
        </q-item-section>
      </q-item>

      <q-item-label header>WeChat</q-item-label>
      <q-item dense>
        <q-item-section avatar>
          <q-icon name="location_city" />
        </q-item-section>
        <q-item-section>
          6666666
        </q-item-section>
      </q-item>
    </q-card>
  </q-dialog>

</template>

<script>
import { defineComponent } from 'vue'
import { get, post, put, deleted } from '../util/request'

export default defineComponent({
  name: 'PageIndex',
  data () {
    return {
      view: false,
      slideVol: 39,
      slideAlarm: 56,
      slideVibration: 63,
      expanded: false,
      lorem: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
    }
  },
  methods: {
    ViewPhone: function (event) {
      this.view = true
    },
    Test () {
      this.$q.notify({
        type: 'warning',
        message: 'This is a "warning" type notification.'
      })
      post('' + this.user.id, {})
        .then(res => {
          this.userGoods = res.data
        })
      deleted('' + this.user.id, {})
        .then(res => {
          this.userGoods = res.data
        })
      put('' + this.user.id, {})
        .then(res => {
          this.userGoods = res.data
        })
    },
    TestAxios () {
      get('/admin/users', {})
        .then(res => {
          console.info(res)
        })
    }
  }
})
</script>

<style lang="sass" scoped>
  .my-card
    width: 100%
    max-width: 350px
</style>
