import Foundation

struct SleepEvent: Identifiable, Decodable {
    let id: UUID
    let type: String
    let timestamp: String
    let level: String

    enum CodingKeys: String, CodingKey {
        case type, timestamp, level
    }

    init(id: UUID = UUID(), type: String, timestamp: String, level: String) {
        self.id = id
        self.type = type
        self.timestamp = timestamp
        self.level = level
    }

    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        id = UUID()
        type = try container.decode(String.self, forKey: .type)
        timestamp = try container.decode(String.self, forKey: .timestamp)
        level = try container.decode(String.self, forKey: .level)
    }
}

struct SleepSummary: Decodable {
    let durationMinutes: Int
    let score: Int
    let advice: String
    let isRecording: Bool

    enum CodingKeys: String, CodingKey {
        case durationMinutes = "duration_m"
        case score, advice
        case isRecording
    }

    init(durationMinutes: Int, score: Int, advice: String, isRecording: Bool) {
        self.durationMinutes = durationMinutes
        self.score = score
        self.advice = advice
        self.isRecording = isRecording
    }

    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        durationMinutes = try container.decode(Int.self, forKey: .durationMinutes)
        score = try container.decode(Int.self, forKey: .score)
        advice = try container.decodeIfPresent(String.self, forKey: .advice) ?? ""
        // isRecording not in backend JSON, default false
        isRecording = false
    }
}
