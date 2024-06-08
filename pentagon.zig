const std = @import("std");

const Point3D = struct {
    x: f64,
    y: f64,
    z: f64,
};

const Point2D = struct {
    x: f64,
    y: f64,
};

const Projection = struct {
    pub fn project(p: Point3D, width: f64, height: f64) Point2D {
        const scale = 20.0;
        return Point2D{
            .x = width / 2.0 + p.x * scale,
            .y = height / 2.0 - p.y * scale,
        };
    }
};

const Rotation = struct {
    pub fn rotateX(p: Point3D, angle: f64) Point3D {
        return Point3D{
            .x = p.x,
            .y = p.y * std.math.cos(angle) - p.z * std.math.sin(angle),
            .z = p.y * std.math.sin(angle) + p.z * std.math.cos(angle),
        };
    }

    pub fn rotateY(p: Point3D, angle: f64) Point3D {
        return Point3D{
            .x = p.x * std.math.cos(angle) + p.z * std.math.sin(angle),
            .y = p.y,
            .z = -p.x * std.math.sin(angle) + p.z * std.math.cos(angle),
        };
    }

    pub fn rotateZ(p: Point3D, angle: f64) Point3D {
        return Point3D{
            .x = p.x * std.math.cos(angle) - p.y * std.math.sin(angle),
            .y = p.x * std.math.sin(angle) + p.y * std.math.cos(angle),
            .z = p.z,
        };
    }
};

fn draw(pentagon: []const Point3D, angle: f64, width: f64, height: f64) void {
  const stdout = std.io.getStdOut().writer();
  var buffer = [_]u8{0} ** (width * height);

  for (buffer_item) |*b| {
    b.* = 32;
  }

  for (pentagon_item) |p| {
    const rotated = Rotation.rotateZ(Rotation.rotateY(Rotation.rotateX(p, angle), angle), angle);
    const projected = Projection.project(rotated, width, height);

    if (projected.x >= 0 and projected.x < width and projected.y >= 0 and projected.y < height) {
      buffer[@intCast(usize, projected.y) * @intCast(usize, width) + @intCast(usize, projected.x)] = 42;
    }
  }

  for (y: usize) |i| {
    stdout.print("{s}\n", .{buffer[y * @intCast(usize, width)..(y + 1) * @intCast(usize, width)]}) catch {};
  }
}

pub fn main() void {
  const width = 80.0;
  const height = 40.0;
  const pentagon = [_]Point3D{
    Point3D{ .x = 0, .y = 1, .z = 0 },
    Point3D{ .x = std.math.cos(2 * std.math.pi / 5), .y = std.math.sin(2 * std.math.pi / 5), .z = 0 },
    Point3D{ .x = std.math.cos(4 * std.math.pi / 5), .y = std.math.sin(4 * std.math.pi / 5), .z = 0 },
    Point3D{ .x = std.math.cos(6 * std.math.pi / 5), .y = std.math.sin(6 * std.math.pi / 5), .z = 0 },
    Point3D{ .x = std.math.cos(8 * std.math.pi / 5), .y = std.math.sin(8 * std.math.pi / 5), .z = 0 },
  };

  var angle: f64 = 0.0;

  while (true) {
    std.os.windowsSleep(100); // Adicione um pequeno delay para visualizar a rotação
    std.os.clearTerminal() catch {};
    draw(pentagon[0..], angle, width, height);
    angle += 0.1;
  }
}
