<template>
  <v-card>
    <v-toolbar flat>
      <v-text-field
        class="mx-4"
        flat
        hide-details
        label="Ricerca per titolo o artista"
        prepend-inner-icon="search"
        v-model="searchString"
        solo-inverted
      ></v-text-field>
    </v-toolbar>
    <v-tabs v-model="tab" background-color="transparent" color="primary" grow>
      <v-tab v-for="item in items" :key="item.text">
        <v-icon>{{ item.icon }} </v-icon>
        <v-chip small v-if="getCounterIfLista(item.component) > 0">{{
          getCounterIfLista(item.component)
        }}</v-chip>
      </v-tab>
    </v-tabs>
    <v-tabs-items v-model="tab">
      <v-tab-item v-for="item in items" :key="item.text">
        <v-card flat color="basil" class="card-overflow">
          <component
            :is="item.component"
            v-bind="getProps(item.component)"
          ></component>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script>
import MusicList from "@/components/MusicList";
import NewMusic from "./NewMusic.vue";
import store from "../store";

export default {
  data() {
    return {
      tab: null,
      items: [
        {
          text: "Nuova canzone",
          component: NewMusic,
          icon: "mdi-plus",
          props: {}
        },
        {
          text: "Lista",
          component: MusicList,
          icon: "mdi-buffer"
        }
      ],
      searchString: ""
    };
  },
  methods: {
    getProps: function(CurrentCoomponent) {
      if (CurrentCoomponent === MusicList) {
        return { searchString: this.searchString };
      }
    },
    getCounterIfLista(currentComponent) {
      if (currentComponent !== MusicList) {
        return;
      }
      const counter = store.getters.musicCountVisible;
      return counter !== 0 ? counter : "";
    }
  },
  watch: {
    searchString: function() {
      if (this.tab !== 1) {
        this.tab = 1;
      }
    }
  }
};
</script>

<style scoped></style>
