<script setup>
import axios from 'axios';

defineProps({
  msg: {
    type: String,
    required: true
  }
})

const API_URL = "http://localhost:8000/down"

const handleDownload = () => {
      axios({
        method: 'post',
        url: API_URL,
        params: {
          issueNo: 168,
        },
        responseType: 'blob'
      }).then(resp => {
        // console.log(resp)
        if (resp.status === 200) {
          alert('生成成功!将自动下载docx文件……');
          let temp = resp.headers["content-disposition"]
            .split(";")[1]
            .split("=")[1];
          console.log(temp)
          let fileName = decodeURIComponent(temp);
          console.log(fileName)
          const blob = new Blob([resp.data]);
          const link = document.createElement("a");

          link.style.display = 'none';
          link.href = URL.createObjectURL(blob);
          link.download = fileName || '趋势周刊word文件.docx'; //下载的文件名
          link.style.display = 'none'

          document.body.appendChild(link);

          link.click();
          document.body.removeChild(link);  // link.remove()
        } else {
          alert(resp.msg);
        }
      });
    }
</script>

<template>
  <div class="greetings">
    <h1 class="green">{{ msg }}</h1>
    <h3>
      You’ve successfully created a project with
      <a href="https://vitejs.dev/" target="_blank" rel="noopener">Vite</a> +
      <a href="https://vuejs.org/" target="_blank" rel="noopener">Vue 3</a>.
    </h3>
    <h3>
      <button @click="handleDownload" style="color: red;">生成并下载.docx</button>
    </h3>
  </div>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}
</style>
