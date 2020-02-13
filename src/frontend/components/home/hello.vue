<template>
  <h1>API says: {{ apiMsg }}</h1>
</template>

<script lang="ts">
import Vue from 'vue'
import HelloRequest from '@/frontend/requests/hello'
import { Component } from 'vue-property-decorator'

@Component
export default class Hello extends Vue {
  getApiMsg = new HelloRequest()
  apiMsg = ''

  created () {
    this.getApiMsg.handle().then(msg => {
      this.apiMsg = msg
    })
  }
}
</script>

<style scoped>

</style>
