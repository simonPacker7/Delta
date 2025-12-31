import Landing from '@/views/landing/Landing.vue'
import Home from '@/views/home/Home.vue'
import Playview from '@/views/play/Play.vue'
import OnlineGameView from '@/views/game/OnlineGame.vue'
import { createRouter, createWebHistory } from 'vue-router'
import { createAuthGuard } from './guards'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: Landing,
    },
    {
      path: '/home',
      component: Home,
    },
    {
      path: '/play',
      children: [
        {
          path: '',
          component: Playview,
        },
        {
          path: 'online',
          component: OnlineGameView,
        }
      ]
    },
  ],
})

router.beforeEach(createAuthGuard());

export default router
