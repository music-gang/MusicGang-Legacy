<template>
  <v-card class="mx-auto" tile full-width>
    <v-list
      :disabled="disabled"
      :dense="dense"
      :two-line="twoLine"
      :three-line="threeLine"
      :shaped="shaped"
      :flat="flat"
      :subheader="subheader"
      :inactive="inactive"
      :sub-group="subGroup"
      :nav="nav"
      :avatar="avatar"
      :rounded="rounded"
    >
      <v-list-item-group color="primary">
        <v-list-item v-for="item in getMusicList" :key="item.id">
          <v-list-item-content class="centered-item-content">
            <v-list-item-title>{{item.title}}</v-list-item-title>
            <v-list-item-subtitle v-if="twoLine || threeLine">{{item.artist}} - {{item.genre}}</v-list-item-subtitle>
            <v-list-item-subtitle v-if="twoLine || threeLine">{{item.notes}}</v-list-item-subtitle>
            <v-list-item-subtitle v-if="twoLine || threeLine">{{item.prompter}}</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </v-card>
</template>

<script>
import store from "../store";
export default {
  name: "MusicList",
  data: () => ({
    disabled: false,
    dense: false,
    twoLine: false,
    threeLine: true,
    shaped: false,
    flat: true,
    subheader: true,
    inactive: false,
    subGroup: false,
    nav: false,
    avatar: false,
    rounded: false
  }),
  computed: {
    getMusicList: () => {
      return store.getters.musicList;
    }
  }
};
</script>

<style scoped>
.centered-item-content {
  text-align: center;
}
</style>