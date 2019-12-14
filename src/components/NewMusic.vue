<template>
  <v-form
    class="mx-auto px-5 form-new-music"
    ref="form"
    v-model="valid"
    lazy-validation
  >
    <v-snackbar
      v-model="show"
      :bottom="y === 'bottom'"
      :color="color"
      :left="x === 'left'"
      :multi-line="mode === 'multi-line'"
      :right="x === 'right'"
      :timeout="timeout"
      :top="y === 'top'"
      :vertical="mode === 'vertical'"
    >
      {{ text }}
      <v-btn dark text @click="snackbar = false">Close</v-btn>
    </v-snackbar>
    <v-text-field
      v-model="tempMusic.title"
      :rules="rules.title"
      label="Titolo"
      required
    ></v-text-field>
    <v-text-field
      v-model="tempMusic.artist"
      :rules="rules.artist"
      label="Artista"
      required
    ></v-text-field>
    <v-text-field
      v-model="tempMusic.genre"
      label="Genere (opzionale)"
    ></v-text-field>
    <v-text-field
      v-model="tempMusic.notes"
      label="Note (opzionale)"
    ></v-text-field>
    <v-text-field
      v-model="tempMusic.prompter"
      label="Suggeritore (opzionale)"
    ></v-text-field>

    <v-row class="justify-center">
      <v-btn :disabled="!valid" color="success" class="mr-4" @click="invia"
        >Invia</v-btn
      >
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
        title: [
          v => {
            return !!v || "Titolo richiesto";
          }
        ],
        artist: [
          v => {
            return !!v || "Artista richiesto";
          }
        ]
      },
      tempMusic: {},
      valid: false,
      color: "success",
      mode: "multi-line",
      snackbar: false,
      text: "Ok!",
      timeout: 1500,
      x: null,
      y: "top",
      show: false
    };
  },
  methods: {
    invia: function() {
      if (
        !this.$refs.form.validate() ||
        this.tempMusic.title === undefined ||
        this.tempMusic.artist === undefined
      ) {
        this.snackbar = true;
        return;
      }
      store.dispatch("addMusic", this.tempMusic).then(() => {
        this.show = true;
        this.$refs.form.reset();
        this.$refs.form.resetValidation();
      });
    }
  }
};
</script>

<style scoped>
.form-new-music {
  height: 100%;
}
</style>
