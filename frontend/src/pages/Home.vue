<template>
  <div class="min-h-screen bg-gradient-to-b from-blue-50 to-white">
    <nav class="sticky top-0 z-40 border-b border-gray-200 bg-white/90 backdrop-blur-md shadow-sm">
      <div class="mx-auto flex h-16 max-w-6xl items-center justify-between px-4">
        <div class="flex items-center gap-3">
          <span class="text-xl font-bold text-blue-600">小智科研</span>
          <span class="text-sm text-gray-500">科学文献综述助手</span>
        </div>
        <div class="flex items-center gap-4">
          <button
            v-if="!authStore.isAuthenticated"
            class="rounded-lg border border-gray-200 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
            @click="authStore.setShowLoginModal(true)"
          >
            登录
          </button>
          <button
            v-else
            class="rounded-lg bg-red-50 px-3 py-2 text-xs font-medium text-red-600 hover:bg-red-100"
            @click="authStore.logout()"
          >
            退出
          </button>
        </div>
      </div>
    </nav>

    <section class="mx-auto max-w-6xl px-4 py-16 text-center">
      <h1 class="mb-6 text-4xl font-bold text-gray-900">
        AI 驱动的文献综述生成平台
      </h1>
      <p class="mx-auto mb-10 max-w-2xl text-lg text-gray-600">
        通过结构化输入快速生成高质量的学术综述，并与后端 Go 服务无缝协作。
      </p>
      <div class="flex flex-col items-center justify-center gap-4 sm:flex-row">
        <button
          class="rounded-lg bg-blue-600 px-6 py-3 text-white hover:bg-blue-700"
          @click="router.push('/ai-writing')"
        >
          一键生成
        </button>
        <button
          class="rounded-lg border border-gray-200 px-6 py-3 text-gray-700 hover:bg-gray-50"
          @click="router.push('/documents')"
        >
          查看文档库
        </button>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";

import { useAuthStore } from "../store/auth";

const router = useRouter();
const authStore = useAuthStore();
</script>
