<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title>
          BBGU二手商品信息平台
        </q-toolbar-title>

        <q-input dark dense standout v-model="text" input-class="text-right" type="search" class="q-ml-md" label="搜索商品">
          <template v-slot:append>
            <q-icon v-if="text === ''" name="search" />
            <q-icon v-else name="clear" class="cursor-pointer" @click="text = ''" />
            <q-btn round dense flat label="search" @click="Search" />
          </template>

        </q-input>

        <q-avatar
          :key="xs"
          :size="lg"
          color="primary"
          text-color="white"
          icon="directions"
        />

        <div class="q-pa-xs" v-if="user !== null">
          <q-btn-dropdown color="primary" :label="user.Username">
            <q-list>
              <q-item clickable v-close-popup @click="onItemClick('/user')">
                <q-item-section>
                  <q-item-label>个人信息</q-item-label>
                </q-item-section>
              </q-item>

              <q-item clickable v-close-popup @click="onItemClick('/userGoods')">
                <q-item-section>
                  <q-item-label>个人商品</q-item-label>
                </q-item-section>
              </q-item>

              <q-item clickable v-close-popup @click="logout">
                <q-item-section>
                  <q-item-label>退出登录</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </q-btn-dropdown>
        </div>

        <div class="q-pa-xs" v-else>
          <q-btn align="left" class="btn-fixed-width push" color="primary" label="登录" to="/login" />
          <q-btn align="right" class="btn-fixed-width push" color="secondary" label="注册" to="/register" />
        </div>

      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
      class="bg-grey-1"
    >
      <q-list>
        <q-item-label
          header
          class="text-grey-8"
        >
          基本操作
        </q-item-label>

        <EssentialLink
          v-for="link in ComputeLinks()"
          :key="link.title"
          v-bind="link"
        />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import EssentialLink from 'components/EssentialLink.vue'

const linksList = [
  {
    title: '管理',
    caption: 'super man',
    icon: 'self_improvement',
    link: '/admin'
  },
  {
    title: '商品列表',
    caption: 'goods list',
    icon: 'storefront',
    link: '/'
  },
  {
    title: '上架商品',
    caption: 'sell',
    icon: 'sell',
    link: '/addGoods'
  },
  {
    title: '个人信息',
    caption: 'user information',
    icon: 'folder_shared',
    link: '/user'
  },
  {
    title: '个人商品',
    caption: 'waitting to build',
    icon: 'cottage',
    link: '/user/goods'
  },
  {
    title: '登录',
    caption: 'login',
    icon: 'how_to_reg',
    link: '/login'
  },
  {
    title: '注册',
    caption: 'register',
    icon: 'person_add_alt',
    link: '/register'
  }
]

import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'MainLayout',

  components: {
    EssentialLink
  },

  data () {
    return {
      user: {},
      text: ''
    }
  },
  created () {
    const localUser = JSON.parse(localStorage.getItem('user'))
    console.info(localUser)
    if (localUser === null) {
      this.user = null
    } else {
      this.user = localUser
    }
  },

  setup () {
    const leftDrawerOpen = ref(false)

    return {
      essentialLinks: linksList,
      leftDrawerOpen,
      toggleLeftDrawer () {
        leftDrawerOpen.value = !leftDrawerOpen.value
      }
    }
  },
  methods: {
    onItemClick (path) {
      this.$router.push(path)
    },
    logout () {
      localStorage.clear()
      this.$router.push('/')
    },
    Search () {
      this.$router.push('?search=' + this.text)
    },
    ComputeLinks () {
      if (this.user !== null && this.user.Role === 'admin') {
        return this.essentialLinks
      } else {
        return this.essentialLinks.slice(1, this.essentialLinks.length)
      }
    }
  }
})
</script>
