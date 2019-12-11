<template>
  <v-app>
    <v-content>
      <Home />
    </v-content>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import Home from './views/Home.vue';
import store from './store';

export default Vue.extend({
  name: 'App',

  components: {
    Home,
  },
  mounted: () => {
    store.dispatch('fetchAll');
  },
});
</script>

<style lang="scss">
$primary: #37c1c3;
$secondary: #f5f5f5;

html,
body,
#app {
  background-color: $secondary;
}

html,
body {
  height: 100%;
}
#app {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  overflow: auto;
}

h1 {
  color: $primary;
}
p {
  color: $primary;
}
a {
  color: $primary;
}

.v-tabs-slider {
  color: $primary;
  caret-color: $primary;
}

.v-tabs--grow > .v-tabs-bar .v-tab {
  background-color: $secondary;
  caret-color: $primary;
  color: $primary;
}

.v-card > :last-child:not(.v-btn):not(.v-chip) {
  background-color: $secondary;
}
</style>