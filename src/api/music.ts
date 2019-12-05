import Axios from "axios"
import { IMusic, Music } from "@/models/music"

interface MusicRequestInterface {
    status: number,
    success: boolean,
    message: string,
    data?: IMusic[]
}

export abstract class MusicApi {

    private static musicAxios = Axios.create();

    static async GetAllMusic(): Promise<Music[]> {

        let url = "/music/all";
        let response = await this.musicAxios.get<MusicRequestInterface>(url);

        let musicList = response.data.data
        if (response.data.status != 0 || !response.data.success || typeof musicList == "undefined") {
            return []
        }

        return musicList.map(musicDto => new Music(musicDto));
    }
}