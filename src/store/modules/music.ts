import {
  Module,
  VuexModule,
  Mutation,
  MutationAction,
  Action
} from "vuex-module-decorators";

import { Music } from "@/models/music";

@Module
export default class MusicModule extends VuexModule {
  list: Array<Music> = [];

  public get musicList(): Array<Music> {
    return this.list;
  }

  @Mutation
  setMusicList(musicList: Array<Music>) {
    this.list = musicList;
  }

  @MutationAction({ mutate: ["list"] })
  async fetchAll() {
    const response = {} as any;
    return response;
  }
}
