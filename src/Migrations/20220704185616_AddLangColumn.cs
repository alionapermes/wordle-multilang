using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace wordlemultilang.Migrations
{
    public partial class AddLangColumn : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AddColumn<string>(
                name: "Language",
                table: "WordOfDayHistory",
                type: "TEXT",
                nullable: false,
                defaultValue: "");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "Language",
                table: "WordOfDayHistory");
        }
    }
}
