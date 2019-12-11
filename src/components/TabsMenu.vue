<template>
  <v-card color="basil">
    <v-tabs v-model="tab" background-color="transparent" color="primary" grow>
      <v-tab v-for="item in items" :key="item.text">
        <v-icon>{{item.icon}}</v-icon>
      </v-tab>
    </v-tabs>
    <v-tabs-items v-model="tab">
      <v-tab-item v-for="item in items" :key="item.text">
        <v-card flat color="basil" class="card-overflow">
          <component :is="item.component"></component>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
import MusicList from "@/components/MusicList";
import NewMusic from "./NewMusic.vue";

export default {
  data() {
    return {
      tab: null,
      items: [
        { text: "Nuova canzone", component: NewMusic, icon: "mdi-plus" },
        { text: "Lista", component: MusicList, icon: "mdi-buffer" }
      ]
    };
  }
};
</script>


<style scoped>
</style>