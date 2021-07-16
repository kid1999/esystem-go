<template>

  <div class="row">
    <div class="col-xs-6 col-sm-4 col-md-3 q-pa-sm" v-for="goods in goodsList" :key="goods.ID">
    <q-card class="shadow-2 "   >
      <q-img
        :src="goods.ImgUrl" height="250px">
        <div class="absolute-bottom text-subtitle text-center">
          {{FormatTime(goods.created_at)}}
        </div>
      </q-img>
      <q-card-section>
        <div class="text-h6">{{ computedGoodsName(goods.GoodsName )}}
          <q-tooltip self="bottom middle" :offset="[10, 10]" anchor="top middle">
            {{goods.GoodsName}}
          </q-tooltip>
        </div>
        <div class="text-body1 text-grey">
          {{goods.Description}}
        </div>
        <div class="text-overline text-orange-9">所有者：{{goods.Username}}</div>
        <div class="row">
          <div class="col-xs-6"><q-btn color="dark" label="联系方式" @click="View(goods)"/></div>
          <div class="col-xs-6"><p class="text-h6 text-grey">¥ {{goods.Price}} 元</p>
          </div>
        </div>

      </q-card-section>
    </q-card>
  </div>
  </div>

  <q-footer style="background-color: transparent;left: 40%;top: 92%;">
    <q-toolbar>
      <q-pagination
        v-model="current"
        :max="totalPage"
        :max-pages="8"
        :boundary-numbers="false"
        direction-links
        push
        color="teal"
      />
    </q-toolbar>
  </q-footer>

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
          {{phone}}
        </q-item-section>
      </q-item>

      <q-item-label header>QQ</q-item-label>
      <q-item dense>
        <q-item-section avatar>
          <q-icon name="support_agent" />
        </q-item-section>
        <q-item-section>
          {{qq}}
        </q-item-section>
      </q-item>

      <q-item-label header>WeChat</q-item-label>
      <q-item dense>
        <q-item-section avatar>
          <q-icon name="textsms" />
        </q-item-section>
        <q-item-section>
          {{wx}}
        </q-item-section>
      </q-item>
    </q-card>
  </q-dialog>

</template>

<script>
import { defineComponent } from 'vue'
import { get } from '../util/request'
import { failMsg } from '../util/msg'
// import { successMsg, failMsg } from '../util/msg'

export default defineComponent({
  name: 'PageIndex',
  data () {
    return {
      view: false,
      current: 1,
      pagesize: 28.0,
      totalPage: 20,
      goodsList: [],
      phone: '',
      wx: '',
      qq: ''
    }
  },
  created () {
    const searchText = this.$route.query.search
    const _this = this
    if (searchText !== undefined) {
      get('/goods/search', { pagesize: this.pagesize, pageindex: this.current, search: searchText }).then(res => {
        _this.goodsList = res.goods
        _this.totalPage = Math.ceil(res.count / _this.pagesize)
        _this.current = res.pageindex
      })
    } else {
      get('/goods', { pagesize: this.pagesize, pageindex: this.current }).then(res => {
        console.info(res)
        _this.goodsList = res.goods
        _this.totalPage = Math.ceil(res.count / _this.pagesize)
        _this.current = res.pageindex
      })
    }
  },
  watch: {
    $route (to, from) { this.$router.go(0) },
    goodsList () {
      if (this.goodsList.length === 0) {
        failMsg('暂无内容被检索！')
      }
    }

  },
  methods: {
    View (goods) {
      this.view = true
      this.qq = goods.QQ
      this.wx = goods.Wx
      this.phone = goods.Phone
    },
    computedGoodsName (str) {
      if (str.length > 12) {
        return str.slice(0, 12) + '...'
      } else {
        return str
      }
    },
    FormatTime (date) {
      return date.slice(0, 10) + ' ' + date.slice(11, 16)
    }
  }
})
</script>
