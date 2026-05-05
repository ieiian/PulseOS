import Foundation

// MARK: - Profile (matches backend user.Profile JSON)

struct ProfileDTO: Codable {
    let id: Int64?
    let name: String
    let age: Int
    let gender: String
    let heightCm: Int
    let weightKg: Int
    let primaryGoal: String
    let secondaryGoals: [String]
    let healthFlags: [String]

    enum CodingKeys: String, CodingKey {
        case id, name, age, gender
        case heightCm = "height_cm"
        case weightKg = "weight_kg"
        case primaryGoal = "primary_goal"
        case secondaryGoals = "secondary_goals"
        case healthFlags = "health_flags"
    }
    
    // Memberwise initializer for creating instances programmatically
    init(id: Int64?, name: String, age: Int, gender: String, heightCm: Int, weightKg: Int, primaryGoal: String, secondaryGoals: [String], healthFlags: [String]) {
        self.id = id
        self.name = name
        self.age = age
        self.gender = gender
        self.heightCm = heightCm
        self.weightKg = weightKg
        self.primaryGoal = primaryGoal
        self.secondaryGoals = secondaryGoals
        self.healthFlags = healthFlags
    }
    
    // Custom decoder for handling null/missing values from API
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        id = try container.decodeIfPresent(Int64.self, forKey: .id)
        name = try container.decodeIfPresent(String.self, forKey: .name) ?? ""
        age = try container.decodeIfPresent(Int.self, forKey: .age) ?? 0
        gender = try container.decodeIfPresent(String.self, forKey: .gender) ?? ""
        heightCm = try container.decodeIfPresent(Int.self, forKey: .heightCm) ?? 0
        weightKg = try container.decodeIfPresent(Int.self, forKey: .weightKg) ?? 0
        primaryGoal = try container.decodeIfPresent(String.self, forKey: .primaryGoal) ?? ""
        secondaryGoals = try container.decodeIfPresent([String].self, forKey: .secondaryGoals) ?? []
        healthFlags = try container.decodeIfPresent([String].self, forKey: .healthFlags) ?? []
    }
}

// MARK: - Settings (matches backend user.Settings JSON)

struct SettingsDTO: Codable {
    let notificationsEnabled: Bool
    let stepPermissionGranted: Bool
    let microphonePermissionGranted: Bool
    let sleepReminderEnabled: Bool

    enum CodingKeys: String, CodingKey {
        case notificationsEnabled = "notifications_enabled"
        case stepPermissionGranted = "step_permission_granted"
        case microphonePermissionGranted = "microphone_permission_granted"
        case sleepReminderEnabled = "sleep_reminder_enabled"
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        notificationsEnabled = try container.decodeIfPresent(Bool.self, forKey: .notificationsEnabled) ?? false
        stepPermissionGranted = try container.decodeIfPresent(Bool.self, forKey: .stepPermissionGranted) ?? false
        microphonePermissionGranted = try container.decodeIfPresent(Bool.self, forKey: .microphonePermissionGranted) ?? false
        sleepReminderEnabled = try container.decodeIfPresent(Bool.self, forKey: .sleepReminderEnabled) ?? false
    }
}

// MARK: - Stats (matches backend user.Stats JSON)

struct StatsDTO: Codable {
    let currentStreak: Int
    let daysTracked: Int

    enum CodingKeys: String, CodingKey {
        case currentStreak = "current_streak"
        case daysTracked = "days_tracked"
    }
    
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        currentStreak = try container.decodeIfPresent(Int.self, forKey: .currentStreak) ?? 0
        daysTracked = try container.decodeIfPresent(Int.self, forKey: .daysTracked) ?? 0
    }
}
