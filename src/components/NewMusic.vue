<template>
  <v-form class="mx-auto px-5 form-new-music" ref="form" v-model="valid" lazy-validation>
    <v-text-field v-model="tempMusic.title" :rules="rules.title" label="Titolo" required></v-text-field>
    <v-text-field v-model="tempMusic.artist" :rules="rules.artist" label="Artista" required></v-text-field>
    <v-text-field v-model="tempMusic.genre" label="Genere (opzionale)"></v-text-field>
    <v-text-field v-model="tempMusic.notes" label="Note (opzionale)"></v-text-field>
    <v-text-field v-model="tempMusic.prompter" label="Suggeritore (opzionale)"></v-text-field>

    <v-row class="justify-center">
      <v-btn :disabled="!valid" color="success" class="mr-4" @click="invia">Invia</v-btn>
    </v-row>
  </v-form>
</template>

<script>
import store from "../store";
export default {
  name: "NewMusic",
  data() {
    return {
      rules: {
        title: [v => !!v || "Titolo richiesto"],
        artist: [v => !!v || "Artista richiesto"]
      },
      tempMusic: {},
      valid: false
    };
  },
  methods: {
    invia: function() {
      if (
        !this.$refs.form.validate() ||
        this.tempMusic.title == undefined ||
        this.tempMusic.artist == undefined
      ) {
        console.log("foffo 2");
        console.log(this.tempMusic);
        console.log(this.$refs.form.validate());
        this.snackbar = true;
        return;
      }
      store.dispatch("addMusic", this.tempMusic);
    }
  }
};
</script>

<style scoped>
.form-new-music {
  height: 100%;
}
</style>