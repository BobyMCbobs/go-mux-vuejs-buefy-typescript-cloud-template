<template>
    <div>
      <section class="hero is-light is-fullheight is-bold">
        <div class="hero-body">
          <div class="container has-text-centered">
            <div class="section">
              <h1 class="title is-1">
                go :heart: vuejs: unknown page
              </h1>
              <h2 class="subtitle is-3">
                <a href="/">[Return home]</a>
              </h2>
            </div>
          </div>
        </div>
      </section>
    </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Component from 'vue-class-component'

@Component({
  name: 'Unknown'
})

export default class Unknown extends Vue {}

</script>
