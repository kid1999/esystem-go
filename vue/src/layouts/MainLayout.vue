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
          BBGU二手交易平台
        </q-toolbar-title>

        <q-input dark dense standout v-model="text" input-class="text-right" class="q-ml-md" label="搜索商品">
          <template v-slot:append>
            <q-icon v-if="text === ''" name="search" />
            <q-icon v-else name="clear" class="cursor-pointer" @click="text = ''" />
          </template>
        </q-input>

        <q-avatar
          :key="xs"
          :size="lg"
          color="primary"
          text-color="white"
          icon="directions"
        />

        <div class="q-pa-xs">
          <q-btn-dropdown color="primary" label="kid1999">
            <q-list>
              <q-item clickable v-close-popup @click="onItemClick">
                <q-item-section>
                  <q-item-label>个人信息</q-item-label>
                </q-item-section>
              </q-item>

              <q-item clickable v-close-popup @click="onItemClick">
                <q-item-section>
                  <q-item-label>个人商品</q-item-label>
                </q-item-section>
              </q-item>

              <q-item clickable v-close-popup @click="onItemClick">
                <q-item-section>
                  <q-item-label>退出登录</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </q-btn-dropdown>
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
          v-for="link in essentialLinks"
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
    title: '个人信息',
    caption: 'quasar.dev',
    icon: 'school',
    link: '#'
  },
  {
    title: '个人商品',
    caption: 'github.com/quasarframework',
    icon: 'code',
    link: '#'
  },
  {
    title: '退出登录',
    caption: 'chat.quasar.dev',
    icon: 'chat',
    link: '#'
  },
  {
    title: '登录',
    caption: 'forum.quasar.dev',
    icon: 'record_voice_over',
    link: '/login'
  },
  {
    title: '注册',
    caption: 'forum.quasar.dev',
    icon: 'record_voice_over',
    link: '/register'
  },
  {
    title: '上架商品',
    caption: 'forum.quasar.dev',
    icon: 'record_voice_over',
    link: '/addGoods'
  }
]

import { defineComponent, ref } from 'vue'

export default defineComponent({
  name: 'MainLayout',

  components: {
    EssentialLink
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
  }
})
</script>
