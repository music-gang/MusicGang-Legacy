export class Music {
  title: string;
  artist: string;
  genre?: string;
  notes?: string;
  prompter?: string;

  constructor(
    title: string,
    artist: string,
    genre?: string,
    notes?: string,
    prompter?: string
  ) {
    this.artist = artist;
    this.title = title;

    this.genre = genre;
    this.notes = notes;
    this.prompter = prompter;
  }
}
