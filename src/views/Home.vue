<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <TabsMenu />
  </div>
</template>

<script>
// @ is an alias to /src
import TabsMenu from '@/components/TabsMenu.vue';

export default {
  name: 'home',
  components: {
    TabsMenu,
  },
};
</script>
