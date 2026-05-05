import Foundation

struct DietPlanOption: Identifiable, Decodable {
    let id: UUID
    let title: String
    let description: String
    let items: [String]

    enum CodingKeys: String, CodingKey {
        case title, description, items
    }

    init(id: UUID = UUID(), title: String, description: String, items: [String]) {
        self.id = id
        self.title = title
        self.description = description
        self.items = items
    }

    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        id = UUID()
        title = try container.decode(String.self, forKey: .title)
        description = try container.decode(String.self, forKey: .description)
        items = try container.decode([String].self, forKey: .items)
    }
}

struct DietStatus: Decodable {
    let recommendation: String
    let summary: String
    let explanation: String

    init(recommendation: String, summary: String, explanation: String) {
        self.recommendation = recommendation
        self.summary = summary
        self.explanation = explanation
    }
}
