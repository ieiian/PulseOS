import Foundation

// MARK: - Today Summary (matches backend meditation.TodaySummary JSON)

struct MeditationTodaySummaryDTO: Codable {
    let totalDurationS: Int
    let completedCount: Int
    let recentSessions: [MeditationSessionDTO]
    let modes: [BreathModeDTO]

    enum CodingKeys: String, CodingKey {
        case totalDurationS = "total_duration_s"
        case completedCount = "completed_count"
        case recentSessions = "recent_sessions"
        case modes
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        totalDurationS = try container.decodeIfPresent(Int.self, forKey: .totalDurationS) ?? 0
        completedCount = try container.decodeIfPresent(Int.self, forKey: .completedCount) ?? 0
        recentSessions = try container.decodeIfPresent([MeditationSessionDTO].self, forKey: .recentSessions) ?? []
        modes = try container.decodeIfPresent([BreathModeDTO].self, forKey: .modes) ?? []
    }
}

struct MeditationSessionDTO: Codable {
    let id: String
    let modeKey: String
    let durationS: Int
    let audioKey: String

    enum CodingKeys: String, CodingKey {
        case id
        case modeKey = "mode_key"
        case durationS = "duration_s"
        case audioKey = "audio_key"
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        id = try container.decodeIfPresent(String.self, forKey: .id) ?? ""
        modeKey = try container.decodeIfPresent(String.self, forKey: .modeKey) ?? ""
        durationS = try container.decodeIfPresent(Int.self, forKey: .durationS) ?? 0
        audioKey = try container.decodeIfPresent(String.self, forKey: .audioKey) ?? ""
    }
}

struct BreathModeDTO: Codable {
    let key: String
    let title: String
    let inhaleSec: Int
    let holdSec: Int
    let exhaleSec: Int
    let description: String

    enum CodingKeys: String, CodingKey {
        case key, title
        case inhaleSec = "inhale_sec"
        case holdSec = "hold_sec"
        case exhaleSec = "exhale_sec"
        case description
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        key = try container.decodeIfPresent(String.self, forKey: .key) ?? ""
        title = try container.decodeIfPresent(String.self, forKey: .title) ?? ""
        inhaleSec = try container.decodeIfPresent(Int.self, forKey: .inhaleSec) ?? 0
        holdSec = try container.decodeIfPresent(Int.self, forKey: .holdSec) ?? 0
        exhaleSec = try container.decodeIfPresent(Int.self, forKey: .exhaleSec) ?? 0
        description = try container.decodeIfPresent(String.self, forKey: .description) ?? ""
    }
}

// MARK: - Session Request

struct MeditationSessionRequestDTO: Codable {
    let modeKey: String
    let durationS: Int
    let audioKey: String

    enum CodingKeys: String, CodingKey {
        case modeKey = "mode_key"
        case durationS = "duration_s"
        case audioKey = "audio_key"
    }
}
