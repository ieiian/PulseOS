import Foundation

// MARK: - Today Summary (matches backend sleep.TodaySummary JSON)

struct SleepTodaySummaryDTO: Codable {
    let session: SleepSessionDTO
    let events: [SleepEventDTO]
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        session = try container.decode(SleepSessionDTO.self, forKey: .session)
        events = try container.decodeIfPresent([SleepEventDTO].self, forKey: .events) ?? []
    }
}

struct SleepSessionDTO: Codable {
    let id: String
    let status: String
    let startedAt: String
    let endedAt: String?
    let durationM: Int
    let score: Int
    let audioURL: String
    let advice: String

    enum CodingKeys: String, CodingKey {
        case id, status
        case startedAt = "started_at"
        case endedAt = "ended_at"
        case durationM = "duration_m"
        case score
        case audioURL = "audio_url"
        case advice
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        id = try container.decodeIfPresent(String.self, forKey: .id) ?? ""
        status = try container.decodeIfPresent(String.self, forKey: .status) ?? ""
        startedAt = try container.decodeIfPresent(String.self, forKey: .startedAt) ?? ""
        endedAt = try container.decodeIfPresent(String.self, forKey: .endedAt)
        durationM = try container.decodeIfPresent(Int.self, forKey: .durationM) ?? 0
        score = try container.decodeIfPresent(Int.self, forKey: .score) ?? 0
        audioURL = try container.decodeIfPresent(String.self, forKey: .audioURL) ?? ""
        advice = try container.decodeIfPresent(String.self, forKey: .advice) ?? ""
    }
}

struct SleepEventDTO: Codable {
    let type: String
    let timestamp: String
    let level: String
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        type = try container.decodeIfPresent(String.self, forKey: .type) ?? ""
        timestamp = try container.decodeIfPresent(String.self, forKey: .timestamp) ?? ""
        level = try container.decodeIfPresent(String.self, forKey: .level) ?? ""
    }
}

// MARK: - Start/End Request

struct SleepStartRequestDTO: Codable {
    let audioURL: String

    enum CodingKeys: String, CodingKey {
        case audioURL = "audio_url"
    }
}

struct SleepEndRequestDTO: Codable {
    let sessionID: String

    enum CodingKeys: String, CodingKey {
        case sessionID = "session_id"
    }
}
