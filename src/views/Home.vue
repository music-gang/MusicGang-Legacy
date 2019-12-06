<template>
  <div class="home">
    <v-row align="center" justify="center">
      <v-img aspect-ratio="1" max-height="250" max-width="250" src="../assets/logo.png" />
    </v-row>
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
