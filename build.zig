const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const mode = b.option(std.builtin.Mode, "mode", "") orelse .Debug;
    const strip = b.option(bool, "strip", "Build without debug info.") orelse false;
    const pie = b.option(bool, "pie", "Build a position independent executable.") orelse false;

    const exe = b.addExecutable(.{
        .name = "slate",
        .root_source_file = .{ .path = "src_zig/main.zig" },
        .target = target,
        .optimize = mode,
    });

    exe.strip = strip;
    exe.pie = pie;
    b.installArtifact(exe);

    const run_cmd = b.addRunArtifact(exe);

    run_cmd.step.dependOn(b.getInstallStep());

    if (b.args) |args| {
        run_cmd.addArgs(args);
    }

    const run_step = b.step("run", "Run the app");
    run_step.dependOn(&run_cmd.step);
}
