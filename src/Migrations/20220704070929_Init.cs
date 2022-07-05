using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace wordlemultilang.Migrations
{
    public partial class Init : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "RussianWords",
                columns: table => new
                {
                    Id = table.Column<ulong>(type: "INTEGER", nullable: false)
                        .Annotation("Sqlite:Autoincrement", true),
                    Text = table.Column<string>(type: "TEXT", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_RussianWords", x => x.Id);
                });

            migrationBuilder.CreateIndex(
                name: "IX_RussianWords_Text",
                table: "RussianWords",
                column: "Text",
                unique: true);
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "RussianWords");
        }
    }
}
