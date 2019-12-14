<template>
  <div class="home">
    <div align="center" justify="center">
      <v-img
        aspect-ratio="1"
        max-height="200"
        max-width="200"
        src="../assets/logo.png"
      />
    </div>
    <TabsMenu />
  </div>
</template>

<script>
import TabsMenu from "@/components/TabsMenu.vue";
export default {
  name: "home",
  components: {
    TabsMenu
  }
};
</script>

<style scoped>
.home {
  overflow: hidden;
}
</style>
