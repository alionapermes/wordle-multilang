using Microsoft.EntityFrameworkCore;

public interface IWordleRepository
{
    public IWordleRepository
    Sync();

    public WordOfDay?
    GetWordOfDay();
}

public class WordleRepository<T> : IWordleRepository where T : BaseWord
{
    private static WordsDbContext _db = new WordsDbContext();

    private string _language;
    private DbSet<T> _words;


    public static IWordleRepository
    GetRussian() => new WordleRepository<RussianWord> {
        _language = "ru",
        _words    = _db.RussianWords
    };

    public static IWordleRepository
    GetEnglish() => new WordleRepository<EnglishWord> {
        _language = "en",
        _words    = _db.EnglishWords
    };

    public IWordleRepository
    Sync()
    {
        var wordOfDay = GetWordOfDay();

        if (wordOfDay == null || wordOfDay.Day < DateTime.Today) {
            wordOfDay = PickWordOfDay();

            _db.Add(wordOfDay);
            _db.SaveChanges();
        }

        return this;
    }

    public WordOfDay?
    GetWordOfDay() => _db.WordOfDayHistory
        .Where(w => w.Language == _language)
        .OrderBy(w => w.Id)
        .LastOrDefault();

    private WordOfDay
    PickWordOfDay()
    {
        var rand      = new Random();
        int wordId    = rand.Next(_words.Count()) + 1;
        BaseWord word = _words
            .Where(w => w.Id == wordId)
            .First();
        
        return new WordOfDay {
            Word     = word.Text,
            Day      = DateTime.Today,
            Language = _language
        };
    }
}

