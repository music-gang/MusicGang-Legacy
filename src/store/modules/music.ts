import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators';

import { IMusic } from '@/models/music';
import { MusicApi } from '@/api/music';

@Module
export default class MusicModule extends VuexModule {
    public list: IMusic[] = [];

    public get musicList(): IMusic[] {
        return this.list;
    }

    @Mutation
    public setMusicList(musicList: IMusic[]) {
        this.list = musicList;
    }

    @Action({ commit: 'setMusicList' })
    public async fetchAll() {
        const response = await MusicApi.GetAllMusic();
        return response;
    }
}
