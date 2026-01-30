<template>
  <section class="mx-auto max-w-4xl px-4 py-12">
    <h2 class="text-2xl font-semibold text-gray-900">AI 文献综述生成</h2>
    <p class="mt-2 text-sm text-gray-600">
      输入标题后将触发 PubMed 检索、摘要整合与文档生成流程。
    </p>

    <form class="mt-8 space-y-4" @submit.prevent="submitWriting">
      <div>
        <label class="text-sm font-medium text-gray-700">综述标题</label>
        <input
          v-model="form.title"
          type="text"
          class="mt-2 w-full rounded-lg border border-gray-200 px-3 py-2 text-sm"
          placeholder="请输入综述标题"
          required
        />
      </div>
      <div>
        <label class="text-sm font-medium text-gray-700">检索主题</label>
        <input
          v-model="form.topic"
          type="text"
          class="mt-2 w-full rounded-lg border border-gray-200 px-3 py-2 text-sm"
          placeholder="例如：阿尔茨海默病治疗"
          required
        />
      </div>
      <div class="flex items-center gap-2 text-sm text-gray-600">
        <input id="includeFilters" v-model="form.includeFilters" type="checkbox" />
        <label for="includeFilters">启用分区/影响因子筛选（将作为二次检索参数）</label>
      </div>
      <button
        type="submit"
        class="rounded-lg bg-blue-600 px-6 py-3 text-white hover:bg-blue-700"
        :disabled="loading"
      >
        {{ loading ? "处理中..." : "生成综述" }}
      </button>
    </form>

    <div v-if="result" class="mt-6 rounded-lg border border-gray-200 bg-white p-4 text-sm">
      <p class="font-medium text-gray-800">状态：{{ result.status }}</p>
      <p class="mt-1 text-gray-600">{{ result.message }}</p>
      <a
        v-if="result.downloadUrl"
        :href="result.downloadUrl"
        class="mt-3 inline-block text-blue-600 underline"
      >
        下载 Word 文档
      </a>
    </div>
  </section>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";

import { writingAPI } from "../services/apiService";

const loading = ref(false);
const result = ref<{ status: string; message: string; downloadUrl?: string } | null>(null);

const form = reactive({
  title: "",
  topic: "",
  includeFilters: false,
});

const submitWriting = async () => {
  loading.value = true;
  result.value = null;
  try {
    const response = await writingAPI.startWriting({
      title: form.title,
      topic: form.topic,
      includeFilters: form.includeFilters,
      locale: "zh-CN",
    });
    result.value = {
      status: response.status ?? "queued",
      message: response.message ?? "已提交生成任务",
      downloadUrl: response.downloadUrl,
    };
  } catch (error) {
    result.value = {
      status: "failed",
      message: "提交失败，请稍后重试。",
    };
  } finally {
    loading.value = false;
  }
};
</script>
