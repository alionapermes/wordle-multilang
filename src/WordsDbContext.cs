using Microsoft.EntityFrameworkCore;


public class WordsDbContext : DbContext
{
    public DbSet<WordOfDay> WordOfDayHistory { get; set; }
    public DbSet<RussianWord> RussianWords { get; set; }
    public DbSet<EnglishWord> EnglishWords { get; set; }
    public string DbPath { get; set; }


    public WordsDbContext()
    {
        var folder =
            Directory.GetParent(Environment.CurrentDirectory).FullName;
        DbPath     = System.IO.Path.Join(folder, "words.db");
    }

    protected override void
    OnConfiguring(DbContextOptionsBuilder opt) =>
        opt.UseSqlite("Data source=../words.db");

    protected override void
    OnModelCreating(ModelBuilder builder)
    {
        builder.Entity<RussianWord>()
            .HasIndex(w => w.Text)
            .IsUnique();
    }
}

