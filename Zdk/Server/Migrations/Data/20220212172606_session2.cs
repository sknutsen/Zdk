using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Zdk.Server.Migrations.Data
{
    public partial class session2 : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropPrimaryKey(
                name: "PK_UserSessions",
                table: "UserSessions");

            migrationBuilder.RenameTable(
                name: "UserSessions",
                newName: "user_sessions");

            migrationBuilder.AddPrimaryKey(
                name: "PK_user_sessions",
                table: "user_sessions",
                column: "user_id");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropPrimaryKey(
                name: "PK_user_sessions",
                table: "user_sessions");

            migrationBuilder.RenameTable(
                name: "user_sessions",
                newName: "UserSessions");

            migrationBuilder.AddPrimaryKey(
                name: "PK_UserSessions",
                table: "UserSessions",
                column: "user_id");
        }
    }
}
