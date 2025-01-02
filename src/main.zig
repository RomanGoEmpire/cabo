const std = @import("std");

const Ability = enum(u2) {
    nothing,
    peek,
    spy,
    swap,
};

const Card = struct { value: u4, ability: Ability };

const Player = struct {
    cards: [4]Card,
    score: u8,

    const default: Player = .{
        .cards = .{ 0, 0, 0, 0 },
        .score = 0,
    };

    fn hand_score(self: Player) u8 {
        var result: u8 = 0;
        for (self.cards) |card| {
            if (card.value > 0) {
                result += card.value - 1;
            }
        }
        return result;
    }
    fn print_player(self: Player) void {
        std.debug.print("Cards\n", .{});
        for (self.cards) |card| {
            std.debug.print("   Card: {}\n", .{card.value - 1});
        }
    }
};

pub fn main() void {
    const player: Player = .{
        .cards = .{
            .{ .value = 2, .ability = Ability.nothing },
            .{ .value = 1, .ability = Ability.nothing },
            .{ .value = 10, .ability = Ability.spy },
            .{ .value = 7, .ability = Ability.nothing },
        },
        .score = 0,
    };

    player.print_player();
    std.debug.print("Hand Score: {}\n", .{player.hand_score()});
}
