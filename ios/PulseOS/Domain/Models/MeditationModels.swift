import Foundation

struct BreathMode: Identifiable, Decodable {
    let id: UUID
    let key: String
    let title: String
    let inhaleSec: Int
    let holdSec: Int
    let exhaleSec: Int
    let description: String

    enum CodingKeys: String, CodingKey {
        case key, title, description
        case inhaleSec = "inhale_sec"
        case holdSec = "hold_sec"
        case exhaleSec = "exhale_sec"
    }

    init(id: UUID = UUID(), key: String, title: String, inhaleSec: Int, holdSec: Int, exhaleSec: Int, description: String) {
        self.id = id
        self.key = key
        self.title = title
        self.inhaleSec = inhaleSec
        self.holdSec = holdSec
        self.exhaleSec = exhaleSec
        self.description = description
    }

    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        id = UUID()
        key = try container.decode(String.self, forKey: .key)
        title = try container.decode(String.self, forKey: .title)
        inhaleSec = try container.decode(Int.self, forKey: .inhaleSec)
        holdSec = try container.decode(Int.self, forKey: .holdSec)
        exhaleSec = try container.decode(Int.self, forKey: .exhaleSec)
        description = try container.decode(String.self, forKey: .description)
    }
}

struct MeditationSummary: Decodable {
    let totalMinutes: Int
    let sessionCount: Int

    enum CodingKeys: String, CodingKey {
        case totalMinutes = "total_duration_s"
        case sessionCount = "completed_count"
    }

    init(totalMinutes: Int, sessionCount: Int) {
        self.totalMinutes = totalMinutes
        self.sessionCount = sessionCount
    }

    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        // Backend returns seconds, convert to minutes
        let totalSeconds = try container.decode(Int.self, forKey: .totalMinutes)
        totalMinutes = totalSeconds / 60
        sessionCount = try container.decode(Int.self, forKey: .sessionCount)
    }
}
