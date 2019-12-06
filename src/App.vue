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

export default Vue.extend({
  name: 'App',

  components: {
    Home,
  },

  data: () => ({
    //
  }),
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