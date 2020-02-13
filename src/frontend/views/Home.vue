<template>
<div>
  <section class="hero is-light is-fullheight is-bold">
    <div class="hero-body">
      <div class="container has-text-centered">
        <div class="section">
          <h1 class="title is-1">
            go :heart: vuejs
          </h1>
          <h2 class="subtitle is-3">
            A cloud-ready {golang, mux} backend + {vuejs, buefy, typescript} frontend
           <br/>
           <hello></hello>
          </h2>
        </div>
      </div>
    </div>
  </section>
</div>
</template>

<script lang="ts">
import hello from '@/frontend/components/home/hello.vue'
import Component from 'vue-class-component'
import Vue from 'vue'

@Component({
  components: {
    hello
  }
})
export default class Home extends Vue {}
</script>

<style scoped>
</style>
