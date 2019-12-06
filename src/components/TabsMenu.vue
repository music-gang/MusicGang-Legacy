<template>
  <v-card color="basil">
    <v-tabs v-model="tab" background-color="transparent" grow>
      <v-tab v-for="item in items" :key="item">{{ item }}</v-tab>
    </v-tabs>
    <v-tabs-items v-model="tab">
      <v-tab-item v-for="item in items" :key="item">
        <v-card flat color="basil" height="max">
          <MusicList v-if="item === 'Nuova canzone'" />
          <MusicList v-if="item === 'Lista'" />
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
import MusicList from "@/components/MusicList";

export default {
  data() {
    return {
      tab: null,
      items: ["Nuova canzone", "Lista"]
    };
  },
  components: {
    MusicList
  }
};
</script>