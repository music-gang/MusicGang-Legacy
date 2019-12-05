import { Module, VuexModule, Mutation, Action } from "vuex-module-decorators";

import { IMusic } from "@/models/music";
import { MusicApi } from "@/api/music";

@Module
export default class MusicModule extends VuexModule {
    list: IMusic[] = [];

    public get musicList(): IMusic[] {
        return this.list;
    }

    @Mutation
    setMusicList(musicList: IMusic[]) {
        this.list = musicList;
    }

    @Action({ commit: "setMusicList" })
    async fetchAll() {
        const response = await MusicApi.GetAllMusic();
        return response;
    }
}
